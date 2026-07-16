package middleware

import (
	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

// InitTracing inicializa o OpenTelemetry com propagador W3C Trace Context.
func InitTracing() {
	tp := sdktrace.NewTracerProvider()
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})
}

// NewRequestIDMiddleware retorna o middleware de tracing distribuído via OpenTelemetry.
func NewRequestIDMiddleware() fiber.Handler {
	return otelfiber.Middleware(
		otelfiber.WithSpanNameFormatter(func(c *fiber.Ctx) string {
			return c.Method() + " " + c.Path()
		}),
	)
}

// TraceID extrai o Trace ID W3C do contexto da requisição.
func TraceID(c *fiber.Ctx) string {
	return trace.SpanFromContext(c.UserContext()).SpanContext().TraceID().String()
}

// SpanID extrai o Span ID W3C do contexto da requisição.
func SpanID(c *fiber.Ctx) string {
	return trace.SpanFromContext(c.UserContext()).SpanContext().SpanID().String()
}

func isHex(s string) bool {
	for _, c := range s {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
			return false
		}
	}
	return true
}
