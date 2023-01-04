package todo

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ovguschin90/todolist/app/todo"
)

type Response struct {
	TaskList *todo.TodoList `json:"task_list"`
}

type Request struct {
	Name string `json:"name"`
}

func List(w http.ResponseWriter, r *http.Request) {
	makeResponse(w)
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	data := r.Body
	defer r.Body.Close()

	body, err := io.ReadAll(data)

	if err != nil {
		http.Error(w, http.ErrBodyNotAllowed.Error(), http.StatusBadRequest)
		return
	}
	var req Request
	_ = json.Unmarshal(body, &req)
	todo.AddTask(req.Name)

	makeResponse(w)
}

func makeResponse(w http.ResponseWriter) {
	var (
		resp []byte
		err  error
	)
	tasks := todo.List()
	if resp, err = json.Marshal(Response{TaskList: tasks}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
