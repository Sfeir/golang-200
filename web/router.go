package web

import (
	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
	"net/http"
)

// Router is the struct use for routing
type Router struct {
	*mux.Router
}

// Route is a structure of Route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// NewRouter creates a new router instance
func NewRouter(controller *TaskController) *Router {
	// new router
	router := Router{mux.NewRouter()}

	// default JSON not found handler
	router.NotFoundHandler = NotFoundHandler()

	// no strict slash
	router.StrictSlash(false)

	// add routes of handler
	for _, route := range controller.Routes {
		logger.WithField("route", route).Debug("adding route to mux")
		// TODO add each route by its Methods, Path (prefix+pattern), Name and handler using the builder
	}

	return &router
}
