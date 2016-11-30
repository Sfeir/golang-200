package web

import (
	"encoding/json"
	"github.com/Sfeir/golang-200/dao"
	"github.com/Sfeir/golang-200/model"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTaskHandlerGet(t *testing.T) {

	// get mock dao
	daoMock, _ := dao.GetTaskDAO("", dao.DAOMock)
	handler := NewTaskHandler(daoMock)

	// build a request
	req, err := http.NewRequest(http.MethodGet, "localhost/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	// build the recorder
	res := httptest.NewRecorder()

	// execute the query
	handler.GetAll(res, req)

	var taskOut []model.Task
	json.NewDecoder(res.Body).Decode(&taskOut)

	if err != nil {
		t.Errorf("Unable to get JSON content %v", err)
	}

	if res.Code != http.StatusOK {
		t.Error("Wrong response code")
	}

	if dao.MockedTask != taskOut[0] {
		t.Errorf("Expected different from %v output %v", dao.MockedTask, taskOut[0])
	}
}

func TestTaskHandlerGetServer(t *testing.T) {

	srv, err := BuildWebServer("", dao.DAOMock, 250*time.Millisecond)

	if err != nil {
		t.Error(err)
	}

	ts := httptest.NewServer(srv)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/tasks")

	if err != nil {
		t.Error(err)
	}

	var resTask []model.Task
	err = json.NewDecoder(res.Body).Decode(&resTask)

	if err != nil {
		t.Errorf("Unable to get JSON content %v", err)
	}

	res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Error("Wrong response code")
	}

	if resTask[0] != dao.MockedTask {
		t.Error("Wrong response body")
	}
}

func BenchmarkTaskHandlerGet(t *testing.B) {

	// tooling purpose
	_ = new([45000]int)

	// get mock dao
	daoMock, _ := dao.GetTaskDAO("", dao.DAOMock)
	handler := NewTaskHandler(daoMock)

	// build a request
	req, err := http.NewRequest("GET", "localhost/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	// build the recorder
	res := httptest.NewRecorder()

	// execute the query
	handler.GetAll(res, req)

	var taskOut []model.Task
	err = json.NewDecoder(res.Body).Decode(&taskOut)

	if err != nil {
		t.Errorf("Unable to get JSON content %v", err)
	}

	expected := model.Task{
		Name:         "Caroni",
		Distiller:    "Caroni",
		Bottler:      "Velier",
		Country:      "Trinidad",
		Composition:  "Molasse",
		SpiritType:   model.TypeRhum,
		Age:          15,
		BottlingDate: time.Date(2015, 01, 01, 0, 0, 0, 0, time.UTC),
		Score:        8.5,
		Comment:      "heavy tire taste",
	}

	if expected != taskOut[0] {
		t.Errorf("Expected different from %v output %v", expected, taskOut)
	}
}
