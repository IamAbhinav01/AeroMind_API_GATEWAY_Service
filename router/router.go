package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Router interface {
	Register(chi.Router)
}

func SetupRouter(UserRouter Router) *chi.Mux {

	chiRouter := chi.NewRouter()
	
	// Basic CORS
	chiRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "x-access-token", "x-idempotency-key"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Root & Health check endpoints for uptime monitors
	chiRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","service":"aeromind-gateway"}`))
	})
	chiRouter.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","service":"aeromind-gateway"}`))
	})

	// Mount application routes under /api/v1 so external clients
	// can call paths like /api/v1/booking/... and match proxy prefixes.
	apiRouter := chi.NewRouter()
	apiRouter.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","service":"aeromind-gateway"}`))
	})
	UserRouter.Register(apiRouter)

	chiRouter.Mount("/api/v1", apiRouter)
	return chiRouter

}