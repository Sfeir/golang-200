package web

import (
	"bytes"
	"encoding/json"
	"github.com/Sfeir/golang-200/dao"
	"github.com/Sfeir/golang-200/model"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// This test is composed by several subtests and uses httptest.ResponseRecorder type to record http response.
// These are NOT end-to-end tests as we are directly calling the handler methods we want to test.
func TestTaskHandlerGet(t *testing.T) {

	// get mock dao
	daoMock, _ := dao.GetTaskDAO("", dao.DAOMock)
	handler := NewTaskHandler(daoMock)

	t.Run("Get all tasks", func(t *testing.T) {
		// build a request
		req := httptest.NewRequest(http.MethodGet, "/tasks", nil)

		// build the recorder
		res := httptest.NewRecorder()

		// execute the query
		handler.GetAll(res, req)

		var taskOut []model.Task
		if err := json.NewDecoder(res.Body).Decode(&taskOut); err != nil {
			t.Errorf("Unable to get JSON content %v", err)
		}

		if res.Code != http.StatusOK {
			t.Errorf("Wrong response code. Got %d instead of %d.", res.Code, http.StatusOK)
		}

		if dao.MockedTask != taskOut[0] {
			t.Errorf("Expected different from %v output %v", dao.MockedTask, taskOut[0])
		}
	})

	t.Run("Get one task", func(t *testing.T) {
		// build a request
		req := httptest.NewRequest(http.MethodGet, "/tasks/"+dao.MockedTask.ID, nil)

		// build the recorder
		res := httptest.NewRecorder()

		// execute the query
		//handler.Get(res, req) ==> the path parameter {id} will not be extracted if you call the handler method directly
		// That's why we use a router instead.
		router := NewRouter(handler)
		router.ServeHTTP(res, req)

		var taskOut model.Task
		if err := json.NewDecoder(res.Body).Decode(&taskOut); err != nil {
			t.Errorf("Unable to get JSON content %v", err)
		}

		if res.Code != http.StatusOK {
			t.Errorf("Wrong response code. Got %d instead of %d.", res.Code, http.StatusOK)
		}

		if dao.MockedTask != taskOut {
			t.Errorf("Expected different from %v output %v", dao.MockedTask, taskOut)
		}
	})

	t.Run("Create a task", func(t *testing.T) {
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

		req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(body))

		// build the recorder
		res := httptest.NewRecorder()

		// execute the query
		handler.Create(res, req)

		var taskOut model.Task
		if err := json.NewDecoder(res.Body).Decode(&taskOut); err != nil {
			t.Errorf("Unable to get JSON content %v", err)
		}

		if res.Code != http.StatusCreated {
			t.Errorf("Wrong response code. Got %d instead of %d.", res.Code, http.StatusCreated)
		}

		task.ID = taskOut.ID
		if task != taskOut {
			t.Errorf("Expected different from %v output %v", task, taskOut)
		}
	})

}

// This test is composed by several subtests and uses httptest.Server type to setup a real local server for testing.
// These ARE end-to-end tests as we are making http requests just as any client would do, without calling any method directly.
func TestTaskHandlerGetServer(t *testing.T) {

	srv, err := BuildWebServer("", dao.DAOMock, 250*time.Millisecond)
	if err != nil {
		t.Error(err)
	}

	ts := httptest.NewServer(srv)
	defer ts.Close()

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

		if resTask[0] != dao.MockedTask {
			t.Error("Wrong response body")
		}
	})

	t.Run("Get one task (end-to-end)", func(t *testing.T) {

		res, err := http.Get(ts.URL + "/tasks/"+dao.MockedTask.ID)
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

		if resTask != dao.MockedTask {
			t.Error("Wrong response body")
		}
	})

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

		res, err := http.Post(ts.URL + "/tasks", "application/json", bytes.NewReader(body))
		if err != nil {
			t.Error(err)
		}

		var taskOut model.Task
		if err := json.NewDecoder(res.Body).Decode(&taskOut); err != nil {
			t.Errorf("Unable to get JSON content %v", err)
		}
		res.Body.Close()

		if res.StatusCode != http.StatusCreated {
			t.Errorf("Wrong response code. Got %d instead of %d.", res.StatusCode, http.StatusCreated)
		}

		task.ID = taskOut.ID
		if task != taskOut {
			t.Errorf("Expected different from %v output %v", task, taskOut)
		}
	})
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
		Title:        "Learn Go",
		Description:  "Let's learn the Go programming language and how to use it in a real project to make great programs.",
		Status:       model.StatusInProgress,
		Priority:     model.PriorityHigh,
		CreationDate: time.Date(2017, 01, 01, 0, 0, 0, 0, time.UTC),
	}

	if expected != taskOut[0] {
		t.Errorf("Expected different from %v output %v", expected, taskOut)
	}
}
