// path: backend/go/fiber/microservices/ms-billing-engine-tax-rates/internal/middleware/metrics.go
package middleware

import (
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

// MetricsCollector armazena métricas agregadas de forma thread-safe.
// Expõe os dados via endpoint /metrics em formato Prometheus text exposition format.
type MetricsCollector struct {
	mu sync.RWMutex

	// http_requests_total
	requestCounter map[string]int64 // "method=POST,path=/calculate,status=200" → count

	// http_request_duration_seconds
	durationBuckets []float64                         // [0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 5]
	durationCount   map[string]int64                  // label → total count
	durationSum     map[string]float64                // label → total sum
	durationBucket  map[string]map[string]int64       // label → bucket → count

	// cache_requests_total
	cacheHits   int64
	cacheMisses int64

	// errors_total
	errorCount map[string]int64 // "type=validation" / "type=internal" → count
}

// NewMetricsCollector cria um coletor de métricas com buckets padrão.
func NewMetricsCollector() *MetricsCollector {
	return &MetricsCollector{
		requestCounter:  make(map[string]int64),
		durationBuckets: []float64{0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 5},
		durationCount:   make(map[string]int64),
		durationSum:     make(map[string]float64),
		durationBucket:  make(map[string]map[string]int64),
		errorCount:      make(map[string]int64),
	}
}

// RecordRequest registra uma requisição HTTP.
func (m *MetricsCollector) RecordRequest(method, path string, status int, duration time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()

	label := "method=" + method + ",path=" + path + ",status=" + strconv.Itoa(status)
	m.requestCounter[label]++

	durLabel := "method=" + method + ",path=" + path
	m.durationCount[durLabel]++
	m.durationSum[durLabel] += duration.Seconds()

	if _, ok := m.durationBucket[durLabel]; !ok {
		m.durationBucket[durLabel] = make(map[string]int64)
	}

	secs := duration.Seconds()
	for _, b := range m.durationBuckets {
		if secs <= b {
			bucketKey := strconv.FormatFloat(b, 'f', -1, 64)
			m.durationBucket[durLabel][bucketKey]++
		}
	}
}

// RecordCacheHit registra um cache hit.
func (m *MetricsCollector) RecordCacheHit() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.cacheHits++
}

// RecordCacheMiss registra um cache miss.
func (m *MetricsCollector) RecordCacheMiss() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.cacheMisses++
}

// RecordError registra um erro.
func (m *MetricsCollector) RecordError(errorType string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.errorCount["type="+errorType]++
}

// Handler retorna o handler Fiber para o endpoint /metrics.
func (m *MetricsCollector) Handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		m.mu.RLock()
		defer m.mu.RUnlock()

		var out string

		// http_requests_total
		out += "# HELP http_requests_total Total de requisições HTTP.\n"
		out += "# TYPE http_requests_total counter\n"
		for label, count := range m.requestCounter {
			out += "http_requests_total{" + label + "} " + strconv.FormatInt(count, 10) + "\n"
		}

		// http_request_duration_seconds
		out += "# HELP http_request_duration_seconds Duração das requisições HTTP.\n"
		out += "# TYPE http_request_duration_seconds histogram\n"
		for label, count := range m.durationCount {
			out += "http_request_duration_seconds_count{" + label + "} " + strconv.FormatInt(count, 10) + "\n"
		}
		for label, sum := range m.durationSum {
			out += "http_request_duration_seconds_sum{" + label + "} " + strconv.FormatFloat(sum, 'f', -1, 64) + "\n"
		}

		// cache_requests_total
		out += "# HELP cache_requests_total Total de requisições de cache.\n"
		out += "# TYPE cache_requests_total counter\n"
		out += "cache_requests_total{result=\"hit\"} " + strconv.FormatInt(m.cacheHits, 10) + "\n"
		out += "cache_requests_total{result=\"miss\"} " + strconv.FormatInt(m.cacheMisses, 10) + "\n"

		// errors_total
		out += "# HELP errors_total Total de erros por tipo.\n"
		out += "# TYPE errors_total counter\n"
		for label, count := range m.errorCount {
			out += "errors_total{" + label + "} " + strconv.FormatInt(count, 10) + "\n"
		}

		c.Set("Content-Type", "text/plain; version=0.0.4")
		return c.SendString(out)
	}
}

// NewMetricsMiddleware cria um middleware Fiber que coleta métricas de requisição.
func NewMetricsMiddleware(collector *MetricsCollector) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		duration := time.Since(start)
		collector.RecordRequest(c.Method(), c.Path(), c.Response().StatusCode(), duration)

		return err
	}
}
