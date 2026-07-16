package middleware

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func testApp() *fiber.App {
	tp := sdktrace.NewTracerProvider()
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	app := fiber.New()
	app.Use(NewRequestIDMiddleware())
	return app
}

func TestMiddleware_WithoutTraceParent(t *testing.T) {
	app := testApp()

	app.Get("/test", func(c *fiber.Ctx) error {
		traceID := TraceID(c)
		spanID := SpanID(c)

		if len(traceID) != 32 || !isHex(traceID) {
			t.Errorf("TraceID invalido: %s", traceID)
		}
		if len(spanID) != 16 || !isHex(spanID) {
			t.Errorf("SpanID invalido: %s", spanID)
		}
		if spanID == "0000000000000000" {
			t.Error("SpanID nao deveria ser zero")
		}

		return c.SendStatus(200)
	})

	req, _ := http.NewRequest("GET", "/test", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("erro no teste: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	}
}

func TestMiddleware_WithValidTraceParent(t *testing.T) {
	app := testApp()

	app.Get("/test", func(c *fiber.Ctx) error {
		traceID := TraceID(c)
		spanID := SpanID(c)

		expectedTraceID := "4bf92f3577b34da6a3ce929d0e0e4736"
		expectedParentSpanID := "00f067aa0ba902b7"

		if traceID != expectedTraceID {
			t.Errorf("TraceID = %s, want %s", traceID, expectedTraceID)
		}
		if spanID == expectedParentSpanID {
			t.Error("SpanID do servico nao deve ser igual ao ParentSpanID")
		}
		if len(spanID) != 16 || !isHex(spanID) {
			t.Errorf("SpanID invalido: %s", spanID)
		}

		return c.SendStatus(200)
	})

	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("traceparent", "00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("erro no teste: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	}
}

func TestMiddleware_WithInvalidTraceParent_GeneratesNew(t *testing.T) {
	app := testApp()

	app.Get("/test", func(c *fiber.Ctx) error {
		traceID := TraceID(c)

		if len(traceID) != 32 || !isHex(traceID) {
			t.Errorf("TraceID deveria ter sido gerado: %s", traceID)
		}

		return c.SendStatus(200)
	})

	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("traceparent", "invalid-header")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("erro no teste: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	}
}

func TestMiddleware_SpanID_Uniqueness(t *testing.T) {
	app := testApp()

	ids := make(map[string]bool)

	app.Get("/test", func(c *fiber.Ctx) error {
		ids[SpanID(c)] = true
		return c.SendStatus(200)
	})

	for i := 0; i < 50; i++ {
		req, _ := http.NewRequest("GET", "/test", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("erro no teste: %v", err)
		}
		if resp.StatusCode != 200 {
			t.Errorf("status = %d, want 200", resp.StatusCode)
		}
	}

	if len(ids) != 50 {
		t.Errorf("esperado 50 IDs unicos, obteve %d", len(ids))
	}
}
