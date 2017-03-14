package dao

import (
	"github.com/Sfeir/golang-200/model"
	"github.com/satori/go.uuid"
	"testing"
	"time"
)

func TestDAOMockInternal(t *testing.T) {

	daoMock := &TaskDAOMock{
		storage: make(map[string]*model.Task),
	}

	toSave := model.Task{
		ID:           uuid.NewV4().String(),
		Title:        "Use Go",
		Description:  "Let's use the Go programming language in a real project.",
		Status:       model.StatusTodo,
		Priority:     model.PriorityMedium,
		CreationDate: time.Date(2017, 02, 01, 0, 0, 0, 0, time.UTC),
		DueDate:      time.Date(2017, 02, 02, 0, 0, 0, 0, time.UTC),
	}

	daoMock.Save(&toSave)

	tasks := daoMock.getBy(func(task *model.Task) bool {
		return task.Status == model.StatusTodo
	})

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
	if tasks[0] != toSave {
		t.Error("Got wrong task from mocked DAO.")
	}

}
