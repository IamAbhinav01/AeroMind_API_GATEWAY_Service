package router

import (

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
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "x-access-token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Mount application routes under /api/v1 so external clients
	// can call paths like /api/v1/booking/... and match proxy prefixes.
	apiRouter := chi.NewRouter()
	UserRouter.Register(apiRouter)

	chiRouter.Mount("/api/v1", apiRouter)
	return chiRouter

}