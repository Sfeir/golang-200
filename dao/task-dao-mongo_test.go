package dao

import (
	"github.com/Sfeir/golang-200/model"
	"os"
	"testing"
	"time"
)

func TestDAOMongo(t *testing.T) {
	// get config
	config := os.Getenv("MONGODB_SRV")

	daoMongo, err := GetTaskDAO(config, DAOMongo)
	if err != nil {
		t.Error(err)
	}

	toSave := model.Task{
		Title:        "Use Go",
		Description:  "Let's use the Go programming language in a real project.",
		Status:       model.StatusTodo,
		Priority:     model.PriorityMedium,
		CreationDate: time.Date(2017, 02, 01, 0, 0, 0, 0, time.UTC),
	}

	err = daoMongo.Save(&toSave)
	if err != nil {
		t.Error(err)
	}

	t.Log("initial task saved", toSave)

	tasks, err := daoMongo.GetAll(NoPaging, NoPaging)
	if err != nil {
		t.Error(err)
	}

	t.Log("initial task found all", tasks[0])

	oneTask, err := daoMongo.GetByID(tasks[0].ID.Hex())
	if err != nil {
		t.Error(err)
	}

	t.Log("initial task found one", oneTask)

	oneTask.Age = 18
	oneTask.Comment = "soft tarmac smell"
	chg, err := daoMongo.Upsert(oneTask.ID.Hex(), oneTask)
	if err != nil {
		t.Error(err)
	}

	t.Log("initial task modified", chg, oneTask)

	oneTask, err = daoMongo.GetByID(oneTask.ID.Hex())
	if err != nil {
		t.Error(err)
	}

	t.Log("initial task found one modified", oneTask)

	err = daoMongo.Delete(oneTask.ID.Hex())
	if err != nil {
		t.Error(err)
	}

	oneTask, err = daoMongo.GetByID(oneTask.ID.Hex())
	if err != nil {
		t.Log("initial task deleted", err, oneTask)
	}

}
