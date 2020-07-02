package routes

import (
	"github.com/gorilla/mux"
)

// Router type
type Router struct {
	R *mux.Router
}

// HandleRoutes : Handles all api routes
func (router *Router) HandleRoutes() *Router {
	// router.R.HandleFunc("/", api.Test)

	return router
}
