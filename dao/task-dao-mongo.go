package dao

import (
	"github.com/Sfeir/golang-200/model"
	"github.com/satori/go.uuid"
	logger "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// compilation time interface check
var _ TaskDAO = (*TaskDAOMongo)(nil)

const (
	collection = "tasks"
	index      = "id"
)

// TaskDAOMongo is the mongo implementation of the TaskDAO
type TaskDAOMongo struct {
	session *mgo.Session
}

// NewTaskDAOMongo creates a new TaskDAO mongo implementation
func NewTaskDAOMongo(session *mgo.Session) TaskDAO {
	// create index
	err := session.DB("").C(collection).EnsureIndex(mgo.Index{
		Key:        []string{index},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	})

	if err != nil {
		logger.WithField("error", err).Warn("mongo db connection")
	}

	return &TaskDAOMongo{
		session: session,
	}
}

// GetByID returns a task by its ID
func (s *TaskDAOMongo) GetByID(ID string) (*model.Task, error) {

	// TODO check in one ligne that the ID is a valid UUID using FromString function
	// TODO fail fast, return an error if its not the case

	// TODO Copy the current MongoDB session
	// TODO defer the close of the session

	// TODO create an empty task to be used as a result container
	// TODO retrieve the collection from the sesison using const
	// TODO Find the unique task having bson.M{"id": ID}
	// TODO return the result and the error
	return nil, nil
}

// getAllTasksByQuery returns tasks by query and paging capability
func (s *TaskDAOMongo) getAllTasksByQuery(query interface{}, start, end int) ([]model.Task, error) {
	// TODO Copy the current MongoDB session
	// TODO defer the close of the session
	// TODO retrieve the collection from the sesison using const

	// TODO check param start and end are not set to NoPaging and are ordered
	hasPaging := true

	// perform request
	var err error
	tasks := []model.Task{}
	if hasPaging {
		// TODO use the Find Method with Skip and Limit (to be calculated) parameters to retrieve all the results
	} else {
		// TODO use only the Find method
	}

	return tasks, err
}

// GetAll returns all tasks with paging capability
func (s *TaskDAOMongo) GetAll(start, end int) ([]model.Task, error) {
	return s.getAllTasksByQuery(nil, start, end)
}

// GetByTitle returns all tasks by title
func (s *TaskDAOMongo) GetByTitle(title string) ([]model.Task, error) {
	return s.getAllTasksByQuery(bson.M{"title": title}, NoPaging, NoPaging)
}

// GetByStatus returns all tasks by status
func (s *TaskDAOMongo) GetByStatus(status int) ([]model.Task, error) {
	return s.getAllTasksByQuery(bson.M{"status": status}, NoPaging, NoPaging)
}

// GetByStatusAndPriority returns all tasks by status and priority
func (s *TaskDAOMongo) GetByStatusAndPriority(status, priority int) ([]model.Task, error) {
	// TODO use the generic getAllTasksByQuery method to build the resultset without paging
	return nil, nil
}

// Save saves the task
func (s *TaskDAOMongo) Save(task *model.Task) error {

	// check task has an ID, if not create one
	if len(task.ID) == 0 {
		task.ID = uuid.NewV4().String()
	}

	session := s.session.Copy()
	defer session.Close()
	c := session.DB("").C(collection)
	return c.Insert(task)
}

// Upsert updates or creates a task, returns true if updated, false otherwise or on error
func (s *TaskDAOMongo) Upsert(task *model.Task) (bool, error) {

	// check ID
	// check task has an ID, if not create one
	if len(task.ID) == 0 {
		task.ID = uuid.NewV4().String()
	}

	session := s.session.Copy()
	defer session.Close()
	c := session.DB("").C(collection)
	chg, err := c.Upsert(bson.M{"id": task.ID}, task)
	if err != nil {
		return false, err
	}
	return chg.Updated > 0, err
}

// Delete deletes a tasks by its ID
func (s *TaskDAOMongo) Delete(ID string) error {

	// TODO check in one ligne that the ID is a valid UUID using FromString function
	// TODO fail fast, return an error if its not the case

	// TODO Copy the current MongoDB session
	// TODO defer the close of the session
	// TODO retrieve the collection from the sesison using const

	// TODO Remove the Task with the given ID

	// TODO return the error
	return nil
}
