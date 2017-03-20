package dao

import (
	"errors"
	"gopkg.in/mgo.v2"
	"time"
)

// DBType lists the type of implementation the factory can return
type DBType int

const (
	// DAOMongo is used for Mongo implementation of TaskDAO
	DAOMongo DBType = iota
	// DAOMock is used for mocked implementation of TaskDAO
	DAOMock

	// mongo timeout
	timeout = 5 * time.Second
	// poolSize of mongo connection pool
	poolSize = 35
)

var (
	// ErrorDAONotFound is used for unknown DAO type
	ErrorDAONotFound = errors.New("Unknown DAO type")
)

// GetTaskDAO returns a TaskDAO according to type and params
func GetTaskDAO(param string, daoType DBType) (TaskDAO, error) {
	switch daoType {
	case DAOMongo:
		// mongo connection
		mgoSession, err := mgo.DialWithTimeout(param, timeout)
		if err != nil {
			return nil, err
		}

		// set 30 sec timeout on session
		mgoSession.SetSyncTimeout(timeout)
		mgoSession.SetSocketTimeout(timeout)
		// set mode
		mgoSession.SetMode(mgo.Monotonic, true)
		// TODO set the connection pool size using the poolSize const

		// TODO return a new DAO Mongo build with the configured session
		return nil, nil
	case DAOMock:
		// TODO return a new mocked DAO instance
		return nil, nil
	default:
		return nil, ErrorDAONotFound
	}
}
