package middleware

import (
	"log"
	"net/http"

	"github.com/felixge/httpsnoop"
	"github.com/go-chi/cors"
)

type Middleware = func(handler http.Handler) http.Handler

func ChainCombine(middlewares ...Middleware) Middleware {
	return func(handler http.Handler) http.Handler {
		for _, m := range middlewares {
			handler = m(handler)
		}
		return handler
	}
}

func WithLogger(handler http.Handler) http.Handler {
	// the creation a handler
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// pass the handler to httpsnoop to get http status and latency
		m := httpsnoop.CaptureMetrics(handler, writer, request)
		// printing extracted data
		log.Printf("http[%d]-- %s -- %s %s\n", m.Code, m.Duration, request.Method, request.URL.Path)
	})
}

func WithCORS(handler http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Use this to allow specific origin hosts, * for matches any origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Signature"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	})
	return c.Handler(handler)
}
