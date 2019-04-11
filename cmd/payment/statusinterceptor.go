package payment

import "net/http"

type statusResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (m *statusResponseWriter) WriteHeader(code int) {
	m.StatusCode = code
	m.ResponseWriter.WriteHeader(code)
}
