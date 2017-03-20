package web

import (
	"github.com/Sfeir/golang-200/dao"
	logger "github.com/Sirupsen/logrus"
	"github.com/meatballhat/negroni-logrus"
	"github.com/urfave/negroni"
	"time"
)

// BuildWebServer constructs a new web server with the right DAO and tasks handler
func BuildWebServer(db string, daoType dao.DBType, statisticsDuration time.Duration) (*negroni.Negroni, error) {

	// TODO fail fast, try to get the DAO of the required type and return (nil,error) if it fails
	// TODO don't forget to log the error and the parameters
	// task dao

	// web server
	n := negroni.New()

	// add middleware for logging
	n.Use(negronilogrus.NewMiddlewareFromLogger(logger.StandardLogger(), "task"))

	// add recovery middleware in case of panic in handler func
	recovery := negroni.NewRecovery()
	recovery.PrintStack = false
	n.Use(recovery)

	// TODO add statistics middleware

	// add as many middleware as you like

	// TODO build a new controller from the DAO

	// TODO build a new router from the controller

	// TODO declare the route handler in last position using UseHandler function

	return n, nil
}
