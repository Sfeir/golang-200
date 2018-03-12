package dao

import (
	"errors"
	"strings"
)

// DBType lists the type of implementation the factory can return
type DBType int

const (
	// DAOMongo is used for Mongo implementation of TaskDAO
	DAOMongo DBType = iota
	// DAOPostgres is used for PostgreSQL implementation of TaskDAO
	DAOPostgres
	// DAOMock is used for mocked implementation of TaskDAO
	DAOMock

	// DAOMockStr is the string representation of the DAOMock DBType
	DAOMockStr = "mock"
)

var (
	// ErrorDAONotFound is used for unknown DAO type
	ErrorDAONotFound = errors.New("unknown DAO type")
)

// ParseDBType parses the string representation and returns the DBType or an error
func ParseDBType(dbType string) (DBType, error) {
	switch strings.ToLower(dbType) {
	case "mongo", "mongodb":
		return DAOMongo, nil
	case "postgre", "postgres", "postgresql":
		return DAOPostgres, nil
	case "mock", "test", "stub", "fake":
		return DAOMock, nil
	}

	return DAOMock, ErrorDAONotFound
}
