package dao_test // /!\ don't change this! create a new test file instead.

import (
	"github.com/Sfeir/golang-200/dao"
	"testing"
)

func TestDAOMock(t *testing.T) {

	daoMock, err := dao.GetTaskDAO("", "", dao.DAOMock)
	if err != nil {
		t.Error(err)
	}

	// TODO create a new task called "toSave" for testing purpose

	// TODO save the "toSave" task
	// TODO check the error

	tasks, err := daoMock.GetAll(dao.NoPaging, dao.NoPaging)
	if err != nil {
		t.Error(err)
	}
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}

	// TODO get "toSave" task by ID and verify that it is successfully retrieved

	// TODO delete the "toSave" task and verify with a get by ID that it is removed

}
