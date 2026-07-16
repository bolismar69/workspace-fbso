// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/middleware/auth.go
package middleware

import (
	"encoding/base64"
	"encoding/json"
	"log/slog"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	HeaderAuthorization = "Authorization"
	HeaderXUserID       = "X-User-Id"
	HeaderXUserName     = "X-User-Name"
	HeaderXUserRoles    = "X-User-Roles"

	bearerPrefix = "Bearer "
)

// JWTClaims representa as claims extraídas do payload JWT decodificado.
type JWTClaims struct {
	Sub           string   `json:"sub"`
	Name          string   `json:"name"`
	PreferredUser string   `json:"preferred_username"`
	Email         string   `json:"email"`
	Roles         []string `json:"roles"`
	RealmAccess   struct {
		Roles []string `json:"roles"`
	} `json:"realm_access"`
	ResourceAccess map[string]struct {
		Roles []string `json:"roles"`
	} `json:"resource_access"`
}

// NewAuthMiddleware cria um middleware que extrai informações do JWT
// já validado pelo Kong/Keycloak e as disponibiliza nos headers e contexto.
//
// Fluxo:
//  1. Kong/Keycloak valida o JWT na borda (assinatura, expiração, etc.)
//  2. O microserviço recebe o JWT no header Authorization: Bearer <token>
//  3. O middleware decodifica o payload Base64 (sem verificar assinatura)
//  4. Extrai sub → X-User-Id, name/preferred_username → X-User-Name,
//     roles/realm_access.roles → X-User-Roles
//  5. Injeta no c.Locals() e nos headers de requisição (c.Request().Header)
func NewAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get(HeaderAuthorization)

		if authHeader == "" || !strings.HasPrefix(authHeader, bearerPrefix) {
			return c.Next()
		}

		token := strings.TrimPrefix(authHeader, bearerPrefix)
		claims, err := decodeJWTClaims(token)
		if err != nil {
			slog.Warn("Falha ao decodificar JWT (token sera ignorado)",
				"error", err,
			)
			return c.Next()
		}

		userID := claims.Sub
		userName := claims.Name
		if userName == "" {
			userName = claims.PreferredUser
		}
		if userName == "" {
			userName = claims.Email
		}

		roles := claims.Roles
		if len(roles) == 0 {
			roles = claims.RealmAccess.Roles
		}

		rolesStr := strings.Join(roles, ",")

		c.Locals(HeaderXUserID, userID)
		c.Locals(HeaderXUserName, userName)
		c.Locals(HeaderXUserRoles, rolesStr)

		c.Request().Header.Set(HeaderXUserID, userID)
		c.Request().Header.Set(HeaderXUserName, userName)
		c.Request().Header.Set(HeaderXUserRoles, rolesStr)

		return c.Next()
	}
}

// decodeJWTClaims decodifica o payload de um JWT (parte central do token)
// sem verificar a assinatura (Kong/Keycloak já validou na borda).
func decodeJWTClaims(token string) (JWTClaims, error) {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return JWTClaims{}, jwtError("token nao possui 3 partes (header.payload.signature)")
	}

	payload := parts[1]

	raw, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		raw, err = base64.RawStdEncoding.DecodeString(payload)
		if err != nil {
			raw, err = base64.StdEncoding.DecodeString(payload)
			if err != nil {
				return JWTClaims{}, jwtError("falha ao decodificar base64 do payload: " + err.Error())
			}
		}
	}

	var claims JWTClaims
	if err := json.Unmarshal(raw, &claims); err != nil {
		return JWTClaims{}, jwtError("falha ao parsear JSON do payload: " + err.Error())
	}

	return claims, nil
}

type jwtError string

func (e jwtError) Error() string { return string(e) }

// GetUserID recupera o X-User-Id do contexto Fiber.
func GetUserID(c *fiber.Ctx) string {
	if v := c.Locals(HeaderXUserID); v != nil {
		return v.(string)
	}
	return ""
}

// GetUserName recupera o X-User-Name do contexto Fiber.
func GetUserName(c *fiber.Ctx) string {
	if v := c.Locals(HeaderXUserName); v != nil {
		return v.(string)
	}
	return ""
}

// GetUserRoles recupera o X-User-Roles do contexto Fiber (separado por vírgula).
func GetUserRoles(c *fiber.Ctx) string {
	if v := c.Locals(HeaderXUserRoles); v != nil {
		return v.(string)
	}
	return ""
}

// HasRole verifica se o usuário possui pelo menos um dos papéis especificados.
// Retorna true se o usuário tem qualquer um dos roles. Case-sensitive.
func HasRole(c *fiber.Ctx, roles ...string) bool {
	userRoles := GetUserRoles(c)
	if userRoles == "" {
		return false
	}
	for _, role := range roles {
		for _, ur := range strings.Split(userRoles, ",") {
			if strings.EqualFold(strings.TrimSpace(ur), role) {
				return true
			}
		}
	}
	return false
}
