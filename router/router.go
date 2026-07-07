package router

import (

	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register(chi.Router)
}

func SetupRouter(UserRouter Router) *chi.Mux {

	chiRouter := chi.NewRouter()

	// Mount application routes under /api/v1 so external clients
	// can call paths like /api/v1/booking/... and match proxy prefixes.
	apiRouter := chi.NewRouter()
	UserRouter.Register(apiRouter)

	chiRouter.Mount("/api/v1", apiRouter)
	return chiRouter

}