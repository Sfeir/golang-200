package dao

import (
	"database/sql"
	// importing postgresql driver for sql connection
	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	// importing file source for db migration
	_ "github.com/mattes/migrate/source/file"
	"gopkg.in/mgo.v2"
	"time"
)

const (
	// db timeout
	timeout = 5 * time.Second

	// poolSize of db connection pool
	poolSize = 35

	// file url scheme
	fileScheme = "file://"
)

// GetTaskDAO returns a TaskDAO according to type and params
func GetTaskDAO(cnxStr, migrationPath string, daoType DBType) (TaskDAO, error) {
	switch daoType {
	case DAOMongo:
		// mongo connection
		mgoSession, err := mgo.DialWithTimeout(cnxStr, timeout)
		if err != nil {
			return nil, err
		}

		// set 30 sec timeout on session
		mgoSession.SetSyncTimeout(timeout)
		mgoSession.SetSocketTimeout(timeout)
		// set mode
		mgoSession.SetMode(mgo.Monotonic, true)
		mgoSession.SetPoolLimit(poolSize)

		return NewTaskDAOMongo(mgoSession), nil
	case DAOPostgres:
		// postgresql connection
		db, err := sql.Open("postgres", cnxStr)

		// check errors
		if err != nil {
			return nil, err
		}

		// set max connection in pool
		// TODO set the maximum open connections in the pool using const

		// try to ping host
		if err = db.Ping(); err != nil {
			return nil, err
		}

		// check is db migration is necessary
		if len(migrationPath) == 0 {
			// TODO if no migration return the new DAO PostgreSQL with the db connection
			return nil, nil
		}

		//  playing database migration
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		m, err := migrate.NewWithDatabaseInstance(
			fileScheme+migrationPath,
			"postgres", driver)

		if err != nil {
			return nil, err
		}

		// upgrade database if necessary
		err = m.Up()
		if err != nil {
			if err != migrate.ErrNoChange {
				return nil, err
			}
		}

		// TODO return the new DAO PostgreSQL with the db connection
		return nil, nil
	case DAOMock:
		return NewTaskDAOMock(), nil
	default:
		return nil, ErrorDAONotFound
	}
}
