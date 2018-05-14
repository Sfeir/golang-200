package dao

import (
	"github.com/Sfeir/golang-200/step05/model"
	"github.com/satori/go.uuid"
	"time"
)

// compilation time interface check
var _ TaskDAO = (*TaskDAOMock)(nil)

// MockedTask is the task returned by this mocked interface
var MockedTask = model.Task{
	ID:           uuid.NewV4().String(),
	Title:        "Learn Go",
	Description:  "Let's learn the Go programming language and how to use it in a real project to make great programs.",
	Status:       model.StatusInProgress,
	Priority:     model.PriorityHigh,
	CreationDate: time.Date(2017, 01, 01, 0, 0, 0, 0, time.UTC),
	DueDate:      time.Date(2017, 01, 02, 0, 0, 0, 0, time.UTC),
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
	task, ok := s.storage[ID]
	if !ok {
		return nil, ErrNotFound
	}
	return task, nil
}

// GetAll returns all tasks with paging capability
func (s *TaskDAOMock) GetAll(start, end int) ([]model.Task, error) {
	if start == NoPaging {
		start = 0
	}
	if end == NoPaging {
		end = len(s.storage) - 1
	}
	if start > end || end > len(s.storage) {
		return []model.Task{}, nil
	}

	tasks, err := s.getBy(func(task *model.Task) bool {
		return true
	})

	// check error
	if err != nil {
		return nil, err
	}

	return tasks[start : end+1], nil
}

// GetByTitle returns all tasks by title
func (s *TaskDAOMock) GetByTitle(title string) ([]model.Task, error) {
	return s.getBy(func(task *model.Task) bool {
		return task.Title == title
	})
}

// GetByStatus returns all tasks by status
func (s *TaskDAOMock) GetByStatus(status model.TaskStatus) ([]model.Task, error) {
	return s.getBy(func(task *model.Task) bool {
		return task.Status == status
	})
}

// GetByStatusAndPriority returns all tasks by status and priority
func (s *TaskDAOMock) GetByStatusAndPriority(status model.TaskStatus, priority model.TaskPriority) ([]model.Task, error) {
	return s.getBy(func(task *model.Task) bool {
		return task.Status == status && task.Priority == priority
	})
}

// getBy returns all tasks that meet filtering conditions
func (s *TaskDAOMock) getBy(filter func(task *model.Task) bool) ([]model.Task, error) {
	// TODO implement the generic filter
	// TODO declare a result array
	// TODO iterate over the storage and apply the filter function
	// TODO check content length and return ErrNotFound if empty
	// TODO return the filtered result

	return nil, nil
}

// Save saves the task
func (s *TaskDAOMock) Save(task *model.Task) error {
	if len(task.ID) == 0 {
		task.ID = uuid.NewV4().String()
	}
	s.storage[task.ID] = task
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
