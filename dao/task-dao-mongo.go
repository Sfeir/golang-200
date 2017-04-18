package dao

import (
	"errors"
	"github.com/Sfeir/golang-200/model"
	logger "github.com/Sirupsen/logrus"
	"github.com/satori/go.uuid"
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

	// check ID
	if _, err := uuid.FromString(ID); err != nil {
		return nil, errors.New("Invalid input to UUID")
	}

	session := s.session.Copy()
	defer session.Close()

	task := model.Task{}
	c := session.DB("").C(collection)
	err := c.Find(bson.M{"id": ID}).One(&task)
	return &task, err
}

// getAllTasksByQuery returns tasks by query and paging capability
func (s *TaskDAOMongo) getAllTasksByQuery(query interface{}, start, end int) ([]model.Task, error) {
	session := s.session.Copy()
	defer session.Close()
	c := session.DB("").C(collection)

	// check param
	hasPaging := start > NoPaging && end > NoPaging && end > start

	// perform request
	var err error
	tasks := []model.Task{}
	if hasPaging {
		err = c.Find(query).Skip(start).Limit(end - start).All(&tasks)
	} else {
		err = c.Find(query).All(&tasks)
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
func (s *TaskDAOMongo) GetByStatus(status model.TaskStatus) ([]model.Task, error) {
	return s.getAllTasksByQuery(bson.M{"status": status}, NoPaging, NoPaging)
}

// GetByStatusAndPriority returns all tasks by status and priority
func (s *TaskDAOMongo) GetByStatusAndPriority(status model.TaskStatus, priority model.TaskPriority) ([]model.Task, error) {
	return s.getAllTasksByQuery(bson.M{"status": status, "priority": priority}, NoPaging, NoPaging)
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

	// check ID
	if _, err := uuid.FromString(ID); err != nil {
		return errors.New("Invalid input to UUID")
	}

	session := s.session.Copy()
	defer session.Close()
	c := session.DB("").C(collection)
	err := c.Remove(bson.M{"id": ID})
	return err
}
