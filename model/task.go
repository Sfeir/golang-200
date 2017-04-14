package model

// TODO build the TaskStatus type as an int
// TaskStatus is the current processing status of a task

// TODO define the Task status enum as const using iota (StatusTodo, StatusInProgress, StatusDone)
const ()

// TODO build the TaskPriority type as an int
// TaskPriority is the priority of a task

// TODO define the Task Priority enum as const using iota (PriorityMinor, PriorityMedium, PriorityHigh)
const ()

// TODO add the Status and Priority enums, the Creation and Due Dates and the JSON ans BSON annotations
// Task is the structure to define a task to be done
type Task struct {
	ID          string
	Title       string
	Description string
	Status      int
	// TODO Priority
	// TODO Creation Date
	// TODO Due Date
}

// TODO add a NewID method to the Task struct to create a new UUID for the task when called
// NewID sets a new ID of the Task as a string
