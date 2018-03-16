package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Sfeir/golang-200/dao"
	"github.com/Sfeir/golang-200/model"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

// This test is composed by several subtests and uses httptest.ResponseRecorder type to record http response.
// These are NOT end-to-end tests as we are directly calling the controller methods we want to test.
func TestTaskControllerGet(t *testing.T) {

	newControllerTest := func() *TaskController {
		daoMock, _ := dao.GetTaskDAO("", "", dao.DAOMock)
		return NewTaskController(daoMock)
	}

	t.Run("Get all tasks", func(t *testing.T) {
		t.Parallel()
		controller := newControllerTest()

		// build a request
		req := httptest.NewRequest(http.MethodGet, "/tasks", nil)

		// build the recorder
		res := httptest.NewRecorder()

		// execute the query
		controller.GetAll(res, req)

		var taskOut []model.Task
		if err := json.NewDecoder(res.Body).Decode(&taskOut); err != nil {
			t.Errorf("Unable to get JSON content %v", err)
		}

		if res.Code != http.StatusOK {
			t.Errorf("Wrong response code. Got %d instead of %d.", res.Code, http.StatusOK)
		}

		if len(taskOut) < 1 {
			t.Fatal("Wrong result size < 1")
		}

		if !dao.MockedTask.Equal(taskOut[0]) {
			t.Errorf("Expected different from %v output %v", dao.MockedTask, taskOut[0])
		}
	})

	t.Run("Get one task", func(t *testing.T) {
		t.Parallel()
		controller := newControllerTest()

		// build a request
		req := httptest.NewRequest(http.MethodGet, "/tasks/"+dao.MockedTask.ID, nil)

		// build the recorder
		res := httptest.NewRecorder()

		// execute the query
		//handler.Get(res, req) ==> the path parameter {id} will not be extracted if you call the handler method directly
		// That's why we use a router instead.
		router := NewRouter(controller)
		router.ServeHTTP(res, req)

		var taskOut model.Task
		if err := json.NewDecoder(res.Body).Decode(&taskOut); err != nil {
			t.Errorf("Unable to get JSON content %v", err)
		}

		if res.Code != http.StatusOK {
			t.Errorf("Wrong response code. Got %d instead of %d.", res.Code, http.StatusOK)
		}

		if !dao.MockedTask.Equal(taskOut) {
			t.Errorf("Expected different from %v output %v", dao.MockedTask, taskOut)
		}
	})

	t.Run("Create a task", func(t *testing.T) {
		t.Parallel()
		//controller := newControllerTest()

		// TODO: implement this unit test to test the creation of a Task through POST /tasks
	})

}

// This test is composed by several subtests and uses httptest.Server type to setup a real local server for testing.
// These ARE end-to-end tests as we are making http requests just as any client would do, without calling any method directly.
func TestTaskControllerGetServer(t *testing.T) {

	// get host IP
	dbHost := os.Getenv("DB_HOST")
	db := fmt.Sprintf("mongodb://%s/tasks", dbHost)

	srv, err := BuildWebServer(db, "", dao.DAOMongo, 250*time.Millisecond)
	if err != nil {
		t.Error(err)
	}

	ts := httptest.NewServer(srv)
	defer ts.Close()

	// task for testing
	var taskTest model.Task

	t.Run("Create a task (end-to-end)", func(t *testing.T) {
		task := model.Task{
			Title:        "Some task",
			Description:  "That's an example of task.",
			Status:       model.StatusTodo,
			Priority:     model.PriorityMinor,
			CreationDate: time.Date(2017, 06, 01, 0, 0, 0, 0, time.UTC),
			DueDate:      time.Date(2017, 07, 12, 0, 0, 0, 0, time.UTC),
		}
		body, _ := json.Marshal(task)

		res, err := http.Post(ts.URL+"/tasks", "application/json", bytes.NewReader(body))
		if err != nil {
			t.Error(err)
		}

		if err := json.NewDecoder(res.Body).Decode(&taskTest); err != nil {
			t.Errorf("Unable to get JSON content %v", err)
		}
		res.Body.Close()

		if res.StatusCode != http.StatusCreated {
			t.Errorf("Wrong response code. Got %d instead of %d.", res.StatusCode, http.StatusCreated)
		}

		task.ID = taskTest.ID
		if !task.Equal(taskTest) {
			t.Errorf("Expected different from %v output %v", task, taskTest)
		}
	})

	t.Run("Get all tasks (end-to-end)", func(t *testing.T) {

		res, err := http.Get(ts.URL + "/tasks")
		if err != nil {
			t.Error(err)
		}

		var resTask []model.Task
		if err := json.NewDecoder(res.Body).Decode(&resTask); err != nil {
			t.Errorf("Unable to get JSON content %v", err)
		}
		res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Errorf("Wrong response code. Got %d instead of %d.", res.StatusCode, http.StatusOK)
		}

		if len(resTask) < 1 {
			t.Fatal("Wrong result size < 1")
		}

		if !resTask[0].Equal(taskTest) {
			t.Errorf("Expected different from %v output %v", resTask[0], taskTest)
		}
	})

	t.Run("Get one task (end-to-end)", func(t *testing.T) {

		res, err := http.Get(ts.URL + "/tasks/" + taskTest.ID)
		if err != nil {
			t.Error(err)
		}

		var resTask model.Task
		if err := json.NewDecoder(res.Body).Decode(&resTask); err != nil {
			t.Errorf("Unable to get JSON content %v", err)
		}
		res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Errorf("Wrong response code. Got %d instead of %d.", res.StatusCode, http.StatusOK)
		}

		if !resTask.Equal(taskTest) {
			t.Errorf("Expected different from %v output %v", resTask, taskTest)
		}
	})

}

func BenchmarkTaskControllerGet(b *testing.B) {

	// get mock dao
	daoMock, _ := dao.GetTaskDAO("", "", dao.DAOMock)
	handler := NewTaskController(daoMock)

	// build a request
	req, err := http.NewRequest(http.MethodGet, "localhost/tasks", nil)
	if err != nil {
		b.Fatal(err)
	}

	for n := 0; n < b.N; n++ {
		// build the recorder
		res := httptest.NewRecorder()

		// execute the query
		handler.GetAll(res, req)
	}
}

func BenchmarkTaskControllerPost(b *testing.B) {

	// get mock dao
	daoMock, _ := dao.GetTaskDAO("", "", dao.DAOMock)
	controller := NewTaskController(daoMock)

	// build a request
	task := model.Task{
		Title:        "Some task",
		Description:  "That's an example of task.",
		Status:       model.StatusTodo,
		Priority:     model.PriorityMinor,
		CreationDate: time.Date(2017, 06, 01, 0, 0, 0, 0, time.UTC),
		DueDate:      time.Date(2017, 07, 12, 0, 0, 0, 0, time.UTC),
	}
	body, _ := json.Marshal(task)

	for n := 0; n < b.N; n++ {
		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(body))
		if err != nil {
			b.Fatal(err)
		}

		// build the recorder
		res := httptest.NewRecorder()

		// execute the query
		controller.Create(res, req)
	}
}

// This benchmark illustrates how memory allocations are visible with PPROF
// $ make benchTool
func BenchmarkHugeMemoryAllocation(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// Make some "huge" memory allocation
		_ = new([45000]int)
	}
}
