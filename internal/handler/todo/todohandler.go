package todo

//	@title			Todo API
//	@version		1.0
//	@description	A simple to-do list API.

//	@tag.name			todos
//	@tag.description	Operations for managing to-do tasks.

//	@accept		json
//	@produce	json

//	@success	200	{array}		todo.Task	"A list of tasks"
//	@failure	400	{object}	api.Error	"Invalid request"
//	@failure	401	{object}	api.Error	"Unauthorized"
//	@failure	500	{object}	api.Error	"Internal server error"

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/ovguschin90/todolist/app/todo"
)

type Response struct {
	TaskList map[uint]*todo.Task `json:"task_list"`
}

type Request struct {
	Name        string `json:"name,omitempty"`
	Due         string `json:"due,omitempty"`
	ID          uint   `json:"id,omitempty"`
	IsCompleted bool   `json:"is_completed"`
}

func (r *Request) GetArray() map[string]string {
	m := make(map[string]string)

	if r.Name != "" {
		m["name"] = r.Name
	}
	if r.Due != "" {
		m["due"] = r.Due
	}
	if r.ID != 0 {
		m["id"] = strconv.Itoa(int(r.ID))
	}
	m["is_completed"] = strconv.FormatBool(r.IsCompleted)

	return m
}

// @router	/tasks [get]
func ListTasks(w http.ResponseWriter, r *http.Request) {
	report := todo.Report()
	makeResponse(w, report)
}

func handleTask(w http.ResponseWriter, r *http.Request, handler func(map[string]string) error) {
	req, err := handleRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler(req.GetArray())
	if err != nil {
		http.Error(w, "bad body", http.StatusBadRequest)
		return
	}

	makeResponse(w, todo.List())
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	handleTask(w, r, todo.AddTask)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	handleTask(w, r, todo.DelTask)
}

func ShowTask(w http.ResponseWriter, r *http.Request) {
	req, err := handleRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := todo.ShowTask(req.GetArray())
	if err != nil {
		http.Error(w, "bad body", http.StatusBadRequest)
		return
	}

	makeResponse(w, task)
}

func EditTask(w http.ResponseWriter, r *http.Request) {
	req, err := handleRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := todo.EditTask(req.GetArray())
	if err != nil {
		http.Error(w, "bad body", http.StatusBadRequest)
		return
	}

	makeResponse(w, task)
}

func makeResponse(w http.ResponseWriter, data interface{}) {
	var (
		resp []byte
		err  error
	)
	if resp, err = json.Marshal(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func handleRequest(r *http.Request) (*Request, error) {
	data := r.Body
	defer r.Body.Close()

	body, err := io.ReadAll(data)
	if err != nil {
		return nil, http.ErrBodyNotAllowed
	}

	req := &Request{}
	err = json.Unmarshal(body, req)
	if err != nil {
		return nil, http.ErrBodyNotAllowed
	}

	err = validate(r, req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func validate(r *http.Request, req *Request) error {
	switch r.URL.Path {
	case Add.String():
		if req.Name == "" {
			return fmt.Errorf("no name")
		}
		if req.Due == "" {
			return fmt.Errorf("no due time")
		}
	case Del.String():
	case Show.String():
	case Edit.String():
		if req.ID == 0 {
			return fmt.Errorf("no id")
		}
	}
	return nil
}
