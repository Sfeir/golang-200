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
		Name:         "Caroni 2000",
		Distiller:    "Caroni",
		Bottler:      "Velier",
		Country:      "Trinidad",
		Composition:  "Melasse",
		SpiritType:   model.TypeRhum,
		Age:          15,
		BottlingDate: time.Date(2015, 01, 01, 0, 0, 0, 0, time.UTC),
		Score:        8.5,
		Comment:      "heavy tire taste",
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
