package todo

import (
	"sync"
)

var (
	TaskList *TodoList
	mu       sync.Mutex
)

type Task struct {
	Name string
	// isCompleted bool
	// Due         *time.Time
}

type TodoList struct {
	tasks []*Task
}

func NewTodoList() {
	TaskList = &TodoList{}
}

func List() *TodoList {
	return TaskList
}

func (tl *TodoList) AddTask(name string) {
	mu.Lock()
	defer mu.Unlock()
	tl.tasks = append(tl.tasks, &Task{
		Name: name,
	})
}

func ShowTask() {

}

func DeleteTask() {

}
