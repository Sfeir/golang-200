package dao

import (
	"github.com/Sfeir/golang-200/model"
	"time"
)

// MockedTask is the task returned by this mocked interface
var MockedTask = model.Task{
	Title:        "Learn Go",
	Description:  "Let's learn the Go programming language and how to use it in a real project to make great programs.",
	Status:       model.StatusInProgress,
	Priority:     model.PriorityHigh,
	CreationDate: time.Date(2017, 01, 01, 0, 0, 0, 0, time.UTC),
	DueDate:      time.Date(2017, 01, 02, 0, 0, 0, 0, time.UTC),
}

// TaskDAOMock is the mocked implementation of the TaskDAO
type TaskDAOMock struct {
}

// NewTaskDAOMock creates a new TaskDAO with a mocked implementation
func NewTaskDAOMock() TaskDAO {
	return &TaskDAOMock{}
}

// GetByID returns a task by its ID
func (s *TaskDAOMock) GetByID(ID string) (*model.Task, error) {
	return &MockedTask, nil
}

// GetAll returns all tasks with paging capability
func (s *TaskDAOMock) GetAll(start, end int) ([]model.Task, error) {
	return []model.Task{MockedTask}, nil
}

// GetByName returns all tasks by name
func (s *TaskDAOMock) GetByName(name string) ([]model.Task, error) {
	return []model.Task{MockedTask}, nil
}

// GetByType returns all tasks by type
func (s *TaskDAOMock) GetByType(taskType string) ([]model.Task, error) {
	return []model.Task{MockedTask}, nil
}

// GetByTypeAndScore returns all tasks by type and score greater than parameter
func (s *TaskDAOMock) GetByTypeAndScore(taskType string, score uint8) ([]model.Task, error) {
	return []model.Task{MockedTask}, nil
}

// Save saves the task
func (s *TaskDAOMock) Save(task *model.Task) error {
	return nil
}

// Upsert updates or creates a task
func (s *TaskDAOMock) Upsert(ID string, task *model.Task) (bool, error) {
	return true, nil
}

// Delete deletes a tasks by its ID
func (s *TaskDAOMock) Delete(ID string) error {
	return nil
}
