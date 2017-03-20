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

	// task dao
	dao, err := dao.GetTaskDAO(db, daoType)
	if err != nil {
		logger.WithField("error", err).WithField("dbtype", dao).WithField("params", db).
			Warn("unable to build the required DAO")
		return nil, err
	}

	// web server
	n := negroni.New()

	// add middleware for logging
	n.Use(negronilogrus.NewMiddlewareFromLogger(logger.StandardLogger(), "task"))

	// add recovery middleware in case of panic in handler func
	recovery := negroni.NewRecovery()
	recovery.PrintStack = false
	n.Use(recovery)

	// add statistics middleware
	n.Use(NewStatisticsMiddleware(statisticsDuration))

	// add as many middleware as you like

	// new controller
	controller := NewTaskController(dao)

	// new router
	router := NewRouter(controller)

	// route handler goes last
	n.UseHandler(router)

	return n, nil
}
