package todo

import (
	"sync"
)

var (
	tasksList *TodoList
	mu        sync.Mutex
)

type Task struct {
	Name string
	// isCompleted bool
	// Due         *time.Time
}

type TodoList struct {
	tasks map[uint]*Task
}

func List() *TodoList {
	initTodoList()
	return tasksList
}

func AddTask(name string) {
	initTodoList()
	mu.Lock()
	defer mu.Unlock()
	tasksList.tasks[getActualIndex(tasksList.tasks)] = &Task{Name: name}
}

func initTodoList() {
	if tasksList == nil {
		tasksList = &TodoList{tasks: make(map[uint]*Task)}
	}
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
