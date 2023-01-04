package todo

import (
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	tasksList   *TodoList
	mu          sync.RWMutex
	timeLayout  = "01/02/2006"
	trimCharset = " "
)

type Task struct {
	Name        string     `json:"name,omitempty"`
	IsCompleted bool       `json:"is_completed,omitempty"`
	Due         *time.Time `json:"due,omitempty"`
}

type TodoList struct {
	tasks map[uint]*Task
}

func List() map[uint]*Task {
	if isEmptyTaskList() {
		initTodoList()
	}
	return tasksList.tasks
}

type ReportResponse struct {
	TaskName string `json:"task_name"`
	Status   string `json:"status"`
}

func Report() []*ReportResponse {
	if isEmptyTaskList() {
		return nil
	}
	mu.RLock()
	defer mu.RUnlock()

	var m []*ReportResponse
	for _, task := range tasksList.tasks {
		r := &ReportResponse{
			TaskName: task.Name,
			Status:   status(task.IsCompleted),
		}

		m = append(m, r)
	}

	return m
}

func status(status bool) string{
	if status {
		return "Done"
	}

	return "In progress"
}

func AddTask(r map[string]string) error {
	if isEmptyTaskList() {
		initTodoList()
	}
	mu.Lock()
	defer mu.Unlock()

	var err error
	if tasksList.tasks[getActualIndex(tasksList.tasks)], err = makeTask(r["name"], r["due"]); err != nil {
		return err
	}

	return nil
}

func DelTask(r map[string]string) error {
	if isEmptyTaskList() {
		return nil
	}

	mu.Lock()
	defer mu.Unlock()

	var (
		err       error
		id        uint
		parsedInt int
	)

	parsedInt, err = strconv.Atoi(r["id"])
	if err != nil {
		return err
	}

	id = uint(parsedInt)
	if _, ok := tasksList.tasks[uint(id)]; ok {
		delete(tasksList.tasks, id)
	}
	return nil
}

func ShowTask(r map[string]string) (*Task, error) {
	if isEmptyTaskList() {
		return nil, nil
	}
	mu.RLock()
	defer mu.RUnlock()

	var (
		err       error
		id        uint
		parsedInt int
	)

	parsedInt, err = strconv.Atoi(r["id"])
	if err != nil {
		return nil, err
	}

	id = uint(parsedInt)
	if _, ok := tasksList.tasks[uint(id)]; ok {
		return tasksList.tasks[uint(id)], nil
	}
	return nil, nil
}

func EditTask(r map[string]string) (*Task, error) {
	if isEmptyTaskList() {
		return nil, nil
	}
	mu.Lock()
	defer mu.Unlock()

	var (
		err       error
		id        uint
		parsedInt int
	)

	parsedInt, err = strconv.Atoi(r["id"])
	if err != nil {
		return nil, err
	}

	id = uint(parsedInt)
	if task, ok := tasksList.tasks[uint(id)]; ok {
		if r["name"] != "" {
			task.Name = r["name"]
		}

		if r["is_completed"] != "" {
			task.IsCompleted, err = strconv.ParseBool(r["is_completed"])
			if err != nil {
				return nil, err
			}
		}

		return task, nil
	}

	return nil, nil
}

func initTodoList() {
	tasksList = &TodoList{tasks: make(map[uint]*Task)}
}

func isEmptyTaskList() bool {
	return tasksList == nil
}

func getLastIndex(list map[uint]*Task) uint {
	var current uint
	if len(list) == 0 {
		return current
	}

	for i := range list {
		if current < i {
			current = i
		}
	}

	return current
}

func getActualIndex(list map[uint]*Task) uint {
	return getLastIndex(list) + 1
}

func makeTask(name, due string) (*Task, error) {
	name = strings.Trim(name, trimCharset)
	due = strings.Trim(due, trimCharset)
	dueDate, err := time.Parse(timeLayout, due)
	if err != nil {
		return nil, err
	}

	return &Task{
		Name: name,
		Due:  &dueDate,
	}, nil
}
