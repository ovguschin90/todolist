package todo

import (
	"encoding/json"
	"net/http"

	"github.com/ovguschin90/todolist/app/todo"
)

type Response struct {
	TaskList *todo.TodoList `json:"task_list"`
}

func List(w http.ResponseWriter, r *http.Request) {
	var (
		resp []byte
		err  error
	)
	if resp, err = json.Marshal(Response{TaskList: todo.List()}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	task := &todo.Task{}
	taskList := todo.List()

	//read POST data
	data := r.PostForm
	task.Name = data.Get("name")
	taskList.AddTask(task.Name)

	var (
		resp []byte
		err  error
	)
	if resp, err = json.Marshal(Response{TaskList: todo.List()}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
