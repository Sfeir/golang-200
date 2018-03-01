package dao_test

import (
	"fmt"
	"github.com/Sfeir/golang-200/dao"
	"github.com/Sfeir/golang-200/model"
	"github.com/satori/go.uuid"
	"os"
	"testing"
	"time"
)

const (
	MIGRATION_PATH = "../etc/migrations"
)

func TestDAOPostgres(t *testing.T) {

	// get host IP
	dbHost := os.Getenv("DB_HOST")
	db := fmt.Sprintf("postgres://todos:password@%s:5432/todos?sslmode=disable&connect_timeout=5", dbHost)

	t.Log(db)

	daoPostgre, err := dao.GetTaskDAO(db, MIGRATION_PATH, dao.DAOPostgres)
	if err != nil {
		t.Error(err)
	}

	title := "Use Go"
	toSave := model.Task{
		ID:           uuid.NewV4().String(),
		Title:        title,
		Description:  "Let's use the Go programming language in a real project.",
		Status:       model.StatusTodo,
		Priority:     model.PriorityMedium,
		CreationDate: time.Date(2017, 02, 01, 0, 0, 0, 0, time.UTC),
		DueDate:      time.Date(2017, 02, 02, 0, 0, 0, 0, time.UTC),
	}

	err = daoPostgre.Save(&toSave)
	if err != nil {
		t.Error(err)
	}

	t.Log("initial task saved", toSave)

	tasks, err := daoPostgre.GetAll(dao.NoPaging, dao.NoPaging)
	if err != nil {
		t.Error(err)
	}

	t.Log("initial task found all", tasks[0])

	oneTask, err := daoPostgre.GetByID(tasks[0].ID)
	if err != nil {
		t.Error(err)
	}

	if !toSave.Equal(*oneTask) {
		t.Error("Failed to save and retrieve task by ID")
	}

	t.Log("initial task found one", oneTask)

	taskByTitle, err := daoPostgre.GetByTitle(title)
	if err != nil {
		t.Error(err)
	}

	if !toSave.Equal(taskByTitle[0]) {
		t.Error("Failed to save and retrieve task by title")
	}

	taskByStatus, err := daoPostgre.GetByStatus(model.StatusTodo)
	if err != nil {
		t.Error(err)
	}

	if !toSave.Equal(taskByStatus[0]) {
		t.Error("Failed to save and retrieve task by status")
	}

	taskByStatusPriority, err := daoPostgre.GetByStatusAndPriority(model.StatusTodo, model.PriorityMedium)
	if err != nil {
		t.Error(err)
	}

	if !toSave.Equal(taskByStatusPriority[0]) {
		t.Error("Failed to save and retrieve task by status and priority")
	}

	oneTask.Title = "Use Go(lang)"
	oneTask.Description = "Let's build a REST service in Go !"
	chg, err := daoPostgre.Upsert(oneTask)
	if err != nil {
		t.Error(err)
	}

	t.Log("initial task modified", chg, oneTask)

	oneTask, err = daoPostgre.GetByID(oneTask.ID)
	if err != nil {
		t.Error(err)
	}

	t.Log("initial task found one modified", oneTask)

	err = daoPostgre.Delete(oneTask.ID)
	if err != nil {
		t.Error(err)
	}

	oneTask, err = daoPostgre.GetByID(oneTask.ID)
	if err != nil {
		t.Log("initial task deleted", err, oneTask)
	}

}
