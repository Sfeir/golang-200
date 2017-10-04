package web

import (
	"github.com/Sfeir/golang-200/dao"
	"github.com/Sfeir/golang-200/model"
	logger "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"net/http"
	"strconv"
)

const (
	prefix = "/tasks"
)

// TaskController is a controller for tasks resource
type TaskController struct {
	taskDao dao.TaskDAO
	Routes  []Route
	Prefix  string
}

// NewTaskController creates a new task controller to manage tasks
func NewTaskController(taskDAO dao.TaskDAO) *TaskController {
	controller := TaskController{
		taskDao: taskDAO,
		Prefix:  prefix,
	}

	// build routes
	routes := []Route{}
	// GetAll
	routes = append(routes, Route{
		Name:        "Get all tasks",
		Method:      http.MethodGet,
		Pattern:     "",
		HandlerFunc: controller.GetAll,
	})
	// Get
	routes = append(routes, Route{
		Name:        "Get one task",
		Method:      http.MethodGet,
		Pattern:     "/{id}",
		HandlerFunc: controller.Get,
	})
	// Create
	routes = append(routes, Route{
		Name:        "Create a task",
		Method:      http.MethodPost,
		Pattern:     "",
		HandlerFunc: controller.Create,
	})
	// Update
	routes = append(routes, Route{
		Name:        "Update a task",
		Method:      http.MethodPut,
		Pattern:     "/{id}",
		HandlerFunc: controller.Update,
	})
	// Delete
	routes = append(routes, Route{
		Name:        "Delete a task",
		Method:      http.MethodDelete,
		Pattern:     "/{id}",
		HandlerFunc: controller.Delete,
	})

	controller.Routes = routes

	return &controller
}

// GetAll retrieve all entities with optional paging of items (start / end are item counts 50 to 100 for example)
func (sh *TaskController) GetAll(w http.ResponseWriter, r *http.Request) {

	startStr := ParamAsString("start", r)
	endStr := ParamAsString("end", r)

	start := dao.NoPaging
	end := dao.NoPaging
	var err error
	if startStr != "" && endStr != "" {
		start, err = strconv.Atoi(startStr)
		if err != nil {
			start = dao.NoPaging
		}
		end, err = strconv.Atoi(endStr)
		if err != nil {
			end = dao.NoPaging
		}
	}

	// find all tasks
	tasks, err := sh.taskDao.GetAll(start, end)
	if err != nil {
		logger.WithField("error", err).Warn("unable to retrieve tasks")
		SendJSONError(w, "Error while retrieving tasks", http.StatusInternalServerError)
		return
	}

	logger.WithField("tasks", tasks).Debug("tasks found")
	SendJSONOk(w, tasks)
}

// Get retrieve an entity by id
func (sh *TaskController) Get(w http.ResponseWriter, r *http.Request) {
	// get the task's ID from the URL
	taskID := ParamAsString("id", r)

	// find the task
	task, err := sh.taskDao.GetByID(taskID)
	if err != nil {
		if err == mgo.ErrNotFound {
			logger.WithField("error", err).WithField("task ID", taskID).Warn("unable to retrieve task by ID")
			SendJSONNotFound(w)
			return
		}

		logger.WithField("error", err).WithField("task ID", taskID).Warn("unable to retrieve task by ID")
		SendJSONError(w, "Error while retrieving task by ID", http.StatusInternalServerError)
		return
	}

	logger.WithField("tasks", task).Debug("task found")
	SendJSONOk(w, task)
}

// Create create an entity
func (sh *TaskController) Create(w http.ResponseWriter, r *http.Request) {
	// task to be created
	task := &model.Task{}
	// get the content body
	err := GetJSONContent(task, r)

	if err != nil {
		logger.WithField("error", err).Warn("unable to decode task to create")
		SendJSONError(w, "Error while decoding task to create", http.StatusBadRequest)
		return
	}

	// save task
	err = sh.taskDao.Save(task)
	if err != nil {
		logger.WithField("error", err).WithField("task", *task).Warn("unable to create task")
		SendJSONError(w, "Error while creating task", http.StatusInternalServerError)
		return
	}

	// send response
	SendJSONWithHTTPCode(w, task, http.StatusCreated)
}

// Update update an entity by id
func (sh *TaskController) Update(w http.ResponseWriter, r *http.Request) {
	// get the task ID from the URL
	taskID := ParamAsString("id", r)

	// task to be created
	task := &model.Task{}
	// get the content body
	err := GetJSONContent(task, r)

	if err != nil {
		logger.WithField("error", err).Warn("unable to decode task to create")
		SendJSONError(w, "Error while decoding task to create", http.StatusBadRequest)
		return
	}

	// save task
	task.ID = taskID
	_, err = sh.taskDao.Upsert(task)
	if err != nil {
		logger.WithField("error", err).WithField("task", *task).Warn("unable to create task")
		SendJSONError(w, "Error while creating task", http.StatusInternalServerError)
		return
	}

	// send response
	SendJSONOk(w, task)
}

// Delete delete an entity by id
func (sh *TaskController) Delete(w http.ResponseWriter, r *http.Request) {
	// get the task ID from the URL
	taskID := ParamAsString("id", r)

	// find task
	err := sh.taskDao.Delete(taskID)
	if err != nil {
		logger.WithField("error", err).WithField("task ID", taskID).Warn("unable to delete task by ID")
		SendJSONError(w, "Error while deleting task by ID", http.StatusInternalServerError)
		return
	}

	logger.WithField("taskID", taskID).Debug("task deleted")
	SendJSONWithHTTPCode(w, nil, http.StatusNoContent)
}
