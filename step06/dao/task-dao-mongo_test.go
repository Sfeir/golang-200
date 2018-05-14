package dao_test

import (
	"fmt"
	"github.com/Sfeir/golang-200/step06/dao"
	"github.com/Sfeir/golang-200/step06/model"
	"github.com/satori/go.uuid"
	"os"
	"testing"
	"time"
)

func TestDAOMongo(t *testing.T) {
	// get host IP
	dbHost := os.Getenv("DB_HOST")
	db := fmt.Sprintf("mongodb://%s/tasks", dbHost)

	daoMongo, err := dao.GetTaskDAO(db, "", dao.DAOMongo)
	if err != nil {
		t.Error(err)
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

	err = daoMongo.Save(&toSave)
	if err != nil {
		t.Error(err)
	}

	t.Log("initial task saved", toSave)

	tasks, err := daoMongo.GetAll(dao.NoPaging, dao.NoPaging)
	if err != nil {
		t.Error(err)
	}

	t.Log("initial task found all", tasks[0])

	oneTask, err := daoMongo.GetByID(tasks[0].ID)
	if err != nil {
		t.Error(err)
	}

	t.Log("initial task found one", oneTask)

	oneTask.Title = "Use Go(lang)"
	oneTask.Description = "Let's build a REST service in Go !"
	chg, err := daoMongo.Upsert(oneTask)
	if err != nil {
		t.Error(err)
	}

	t.Log("initial task modified", chg, oneTask)

	oneTask, err = daoMongo.GetByID(oneTask.ID)
	if err != nil {
		t.Error(err)
	}

	t.Log("initial task found one modified", oneTask)

	err = daoMongo.Delete(oneTask.ID)
	if err != nil {
		t.Error(err)
	}

	oneTask, err = daoMongo.GetByID(oneTask.ID)
	if err != nil {
		t.Log("initial task deleted", err, oneTask)
	}

}
