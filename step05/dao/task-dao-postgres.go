package dao

import (
	"database/sql"
	"fmt"
	"github.com/Sfeir/golang-200/step05/model"
	"github.com/satori/go.uuid"
)

// compilation time interface check
var _ TaskDAO = (*TaskDAOPostgres)(nil)

// TaskDAOPostgres is the postgreSQL implementation of the TaskDAO
type TaskDAOPostgres struct {
	db *sql.DB
}

// NewTaskDAOPostgres creates a new TaskDAO postgreSQL implementation
func NewTaskDAOPostgres(db *sql.DB) TaskDAO {
	return &TaskDAOPostgres{
		db: db,
	}
}

// GetByID returns a task by its ID
func (s *TaskDAOPostgres) GetByID(ID string) (*model.Task, error) {

	// TODO check ID is a valid UUID return ErrInvalidUUID if necessary

	// TODO use Query method to request the task with the given ID

	// TODO if result is OK, map the row to a Task object using the mapRows function

	// TODO if results length is 0, return a ErrNotFound

	// TODO if result length is > 0, return a new error "too many results for UUID + ID"

	// TODO finally if ok return the only result

	return nil, nil
}

// GetAll returns all tasks with paging capability
func (s *TaskDAOPostgres) GetAll(start, end int) ([]model.Task, error) {

	// check param
	hasPaging := start > NoPaging && end > NoPaging && end >= start

	query := `SELECT * FROM todos`
	if hasPaging {
		pagingQuery := ` OFFSET %d LIMIT %d`
		query = fmt.Sprintf(query+pagingQuery, start, end-start+1)
	}

	// query db
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	// map results to Task slice
	return mapRows(rows)
}

// GetByTitle returns all tasks by title
func (s *TaskDAOPostgres) GetByTitle(title string) ([]model.Task, error) {
	query := `SELECT * FROM todos WHERE title LIKE '%' || $1 || '%'`

	// query db
	rows, err := s.db.Query(query, title)
	if err != nil {
		return nil, err
	}

	// map results to Task slice
	return mapRows(rows)
}

// GetByStatus returns all tasks by status
func (s *TaskDAOPostgres) GetByStatus(status model.TaskStatus) ([]model.Task, error) {
	query := `SELECT * FROM todos WHERE status=$1`

	// query db
	rows, err := s.db.Query(query, status)
	if err != nil {
		return nil, err
	}

	// map results to Task slice
	return mapRows(rows)
}

// GetByStatusAndPriority returns all tasks by status and priority
func (s *TaskDAOPostgres) GetByStatusAndPriority(status model.TaskStatus, priority model.TaskPriority) ([]model.Task, error) {
	query := `SELECT * FROM todos WHERE status=$1 AND priority=$2`

	// query db
	rows, err := s.db.Query(query, status, priority)
	if err != nil {
		return nil, err
	}

	// map results to Task slice
	return mapRows(rows)
}

// Save saves the task
func (s *TaskDAOPostgres) Save(task *model.Task) error {

	// check task has an ID, if not create one
	if len(task.ID) == 0 {
		task.ID = uuid.NewV4().String()
	}

	query := `INSERT INTO todos(uuid,title,description,status,priority,creation_date,due_date) VALUES($1,$2,$3,$4,$5,$6,$7)`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(task.ID, task.Title, task.Description, task.Status, task.Priority, task.CreationDate, task.DueDate)

	return err
}

// Upsert updates or creates a task, returns true if updated, false otherwise or on error
func (s *TaskDAOPostgres) Upsert(task *model.Task) (bool, error) {

	// check task has an ID, if not create the task
	if len(task.ID) == 0 {
		return false, s.Save(task)
	}

	_, err := s.GetByID(task.ID)
	if err != nil && err != ErrNotFound {
		return false, err
	}

	query := `UPDATE todos SET title=$2,description=$3,status=$4,priority=$5,creation_date=$6,due_date=$7 WHERE uuid=$1`
	res, err := s.db.Exec(query, task.ID, task.Title, task.Description, task.Status, task.Priority, task.CreationDate, task.DueDate)
	if err != nil {
		return false, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return count == 1, nil
}

// Delete deletes a tasks by its ID
func (s *TaskDAOPostgres) Delete(ID string) error {

	// check ID
	if _, err := uuid.FromString(ID); err != nil {
		return ErrInvalidUUID
	}

	query := `DELETE FROM todos WHERE uuid=$1`
	_, err := s.db.Exec(query, ID)
	if err != nil {
		return err
	}

	return nil
}

// mapRows maps a rows with their columns to a slice of Task
func mapRows(rows *sql.Rows) ([]model.Task, error) {
	// hydrate results
	var tasks []model.Task
	for rows.Next() {
		task := model.Task{}
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Priority, &task.CreationDate, &task.DueDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
