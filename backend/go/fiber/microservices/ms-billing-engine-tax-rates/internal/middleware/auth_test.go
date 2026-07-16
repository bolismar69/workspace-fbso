package middleware

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestDecodeJWTClaims_Valid(t *testing.T) {
	// Header: {"alg":"RS256","typ":"JWT"}
	// Payload: {"sub":"user-123","name":"John Doe","preferred_username":"johndoe","roles":["admin","tax-analyst"]}
	// Signature: ...
	const validToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ1c2VyLTEyMyIsIm5hbWUiOiJKb2huIERvZSIsInByZWZlcnJlZF91c2VybmFtZSI6ImpvaG5kb2UiLCJyb2xlcyI6WyJhZG1pbiIsInRheC1hbmFseXN0Il19.signature"

	claims, err := decodeJWTClaims(validToken)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	if claims.Sub != "user-123" {
		t.Errorf("sub = %s, want user-123", claims.Sub)
	}
	if claims.Name != "John Doe" {
		t.Errorf("name = %s, want John Doe", claims.Name)
	}
	if claims.PreferredUser != "johndoe" {
		t.Errorf("preferred_username = %s, want johndoe", claims.PreferredUser)
	}
	if len(claims.Roles) != 2 || claims.Roles[0] != "admin" {
		t.Errorf("roles = %v, want [admin tax-analyst]", claims.Roles)
	}
}

func TestDecodeJWTClaims_RealmRoles(t *testing.T) {
	const token = "eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJ1c2VyLTQ1NiIsImVtYWlsIjoiYWRtaW5AdGF4bmV4dXMuY29tIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9wZXJhdG9yIiwidmlld2VyIl19fQ.sig"

	claims, err := decodeJWTClaims(token)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	if claims.Sub != "user-456" {
		t.Errorf("sub = %s, want user-456", claims.Sub)
	}
	if len(claims.RealmAccess.Roles) != 2 || claims.RealmAccess.Roles[0] != "operator" {
		t.Errorf("realm_access.roles = %v, want [operator viewer]", claims.RealmAccess.Roles)
	}
}

func TestDecodeJWTClaims_PaddedBase64(t *testing.T) {
	const token = "header.eyJzdWIiOiJwYWRkZWQifQ==.sig"

	claims, err := decodeJWTClaims(token)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	if claims.Sub != "padded" {
		t.Errorf("sub = %s, want padded", claims.Sub)
	}
}

func TestDecodeJWTClaims_Invalid(t *testing.T) {
	tests := []struct {
		name  string
		token string
	}{
		{"vazio", ""},
		{"sem dots", "abc123"},
		{"apenas header", "header."},
		{"header e dot", "header.payload"},
		{"base64 invalido", "a.b$$$.c"},
		{"json invalido", "a.eyJub3QtanNvbn0.c"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := decodeJWTClaims(tt.token)
			if err == nil {
				t.Errorf("decodeJWTClaims(%q) deveria ter falhado", tt.token)
			}
		})
	}
}

func TestAuthMiddleware_WithValidToken(t *testing.T) {
	app := fiber.New()
	app.Use(NewAuthMiddleware())

	app.Get("/test", func(c *fiber.Ctx) error {
		userID := GetUserID(c)
		userName := GetUserName(c)
		userRoles := GetUserRoles(c)

		if userID != "user-123" {
			t.Errorf("userID = %s, want user-123", userID)
		}
		if userName != "John Doe" {
			t.Errorf("userName = %s, want John Doe", userName)
		}
		if userRoles != "admin,tax-analyst" {
			t.Errorf("userRoles = %s, want admin,tax-analyst", userRoles)
		}

		return c.SendStatus(200)
	})

	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set(HeaderAuthorization, "Bearer eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJ1c2VyLTEyMyIsIm5hbWUiOiJKb2huIERvZSIsInByZWZlcnJlZF91c2VybmFtZSI6ImpvaG5kb2UiLCJyb2xlcyI6WyJhZG1pbiIsInRheC1hbmFseXN0Il19.sig")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("erro no teste: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	}
}

func TestAuthMiddleware_WithoutToken(t *testing.T) {
	app := fiber.New()
	app.Use(NewAuthMiddleware())

	app.Get("/test", func(c *fiber.Ctx) error {
		if GetUserID(c) != "" {
			t.Error("userID deveria estar vazio sem token")
		}
		return c.SendStatus(200)
	})

	req, _ := http.NewRequest("GET", "/test", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("erro no teste: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("status = %d, want 200 (continua sem auth)", resp.StatusCode)
	}
}

func TestAuthMiddleware_WithInvalidToken(t *testing.T) {
	app := fiber.New()
	app.Use(NewAuthMiddleware())

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set(HeaderAuthorization, "Bearer invalid.token.here")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("erro no teste: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("status = %d, want 200 (continua mesmo com token invalido)", resp.StatusCode)
	}
}

func TestAuthMiddleware_WithRealmRoles(t *testing.T) {
	app := fiber.New()
	app.Use(NewAuthMiddleware())

	app.Get("/test", func(c *fiber.Ctx) error {
		roles := GetUserRoles(c)
		if roles != "operator,viewer" {
			t.Errorf("roles = %s, want operator,viewer", roles)
		}
		return c.SendStatus(200)
	})

	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set(HeaderAuthorization, "Bearer eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJ1c2VyLTQ1NiIsImVtYWlsIjoiYWRtaW5AdGF4bmV4dXMuY29tIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9wZXJhdG9yIiwidmlld2VyIl19fQ.sig")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("erro no teste: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	}
}

func TestAuthMiddleware_SetsHeaders(t *testing.T) {
	app := fiber.New()
	app.Use(NewAuthMiddleware())

	app.Get("/test", func(c *fiber.Ctx) error {
		if c.Get(HeaderXUserID) != "user-123" {
			t.Errorf("X-User-Id header: %s", c.Get(HeaderXUserID))
		}
		return c.SendStatus(200)
	})

	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set(HeaderAuthorization, "Bearer eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJ1c2VyLTEyMyIsIm5hbWUiOiJKb2huIERvZSIsInJvbGVzIjpbXX0.sig")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("erro no teste: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	}
}
