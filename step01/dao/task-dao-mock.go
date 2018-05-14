package dao

import (
	"github.com/Sfeir/golang-200/step01/model"
	"github.com/satori/go.uuid"
)

// compilation time interface check
var _ TaskDAO = (*TaskDAOMock)(nil)

// TODO add the missing attributes (ID, Priority, dates) to have a complete Task

// MockedTask is the task returned by this mocked interface
var MockedTask = model.Task{
	Title:       "Learn Go",
	Description: "Let's learn the Go programming language and how to use it in a real project to make great programs.",
	Status:      0,
}

// TaskDAOMock is the mocked implementation of the TaskDAO
type TaskDAOMock struct {
	storage map[string]*model.Task
}

// NewTaskDAOMock creates a new TaskDAO with a mocked implementation
func NewTaskDAOMock() TaskDAO {
	daoMock := &TaskDAOMock{
		storage: make(map[string]*model.Task),
	}

	// Adds some fake data
	daoMock.Save(&MockedTask)

	return daoMock
}

// GetByID returns a task by its ID
func (s *TaskDAOMock) GetByID(ID string) (*model.Task, error) {
	// TODO check the map contains the ID

	// TODO if not return nil and a new error

	// TODO if ok return the pointer to the task
	return nil, nil
}

// GetAll returns all tasks with paging capability
func (s *TaskDAOMock) GetAll(start, end int) ([]model.Task, error) {
	if start == NoPaging {
		start = 0
	}
	if end == NoPaging {
		end = len(s.storage)
	}
	if start > end || end > len(s.storage) {
		return []model.Task{}, nil
	}

	tasks := s.getBy(func(task *model.Task) bool {
		return true
	})

	return tasks[start:end], nil
}

// GetByTitle returns all tasks by title
func (s *TaskDAOMock) GetByTitle(title string) ([]model.Task, error) {
	tasks := s.getBy(func(task *model.Task) bool {
		return task.Title == title
	})
	return tasks, nil
}

// GetByStatus returns all tasks by status
func (s *TaskDAOMock) GetByStatus(status int) ([]model.Task, error) {
	// TODO implement the GetByStatus function using an anonymous function comparing the status to the param
	return nil, nil
}

// GetByStatusAndPriority returns all tasks by status and priority
func (s *TaskDAOMock) GetByStatusAndPriority(status, priority int) ([]model.Task, error) {
	tasks := s.getBy(func(task *model.Task) bool {
		// TODO implement the status and priority filter
		return true
	})
	return tasks, nil
}

// getBy returns all tasks that meet filtering conditions
func (s *TaskDAOMock) getBy(filter func(task *model.Task) bool) []model.Task {
	var tasks []model.Task
	for _, task := range s.storage {
		if filter(task) {
			tasks = append(tasks, *task)
		}
	}
	return tasks
}

// Save saves the task
func (s *TaskDAOMock) Save(task *model.Task) error {
	// TODO check that the task has an ID
	// TODO if not add one

	// TODO save the newly created task to the map
	return nil
}

// Upsert updates or creates a task
func (s *TaskDAOMock) Upsert(task *model.Task) (bool, error) {
	// check ID
	if len(task.ID) == 0 {
		task.ID = uuid.NewV4().String()
	}

	s.Save(task)
	return true, nil
}

// Delete deletes a tasks by its ID
func (s *TaskDAOMock) Delete(ID string) error {
	delete(s.storage, ID)
	return nil
}
