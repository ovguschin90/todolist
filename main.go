package main

import (
	"net/http"

	h "github.com/ovguschin90/todolist/internal/handler"
	th "github.com/ovguschin90/todolist/internal/handler/todo"
	"github.com/ovguschin90/todolist/internal/router"
)

func main() {
	router := &router.Router{}

	//AddRoutes
	router.AddRoute(http.MethodGet, "/", h.Index)
	router.AddRoute(http.MethodGet, "/todos", th.List)
	// router.AddRoute("/todos/id=[0-9]+", todo.ShowTask)
	router.AddRoute(http.MethodPost, "/todos/add", th.AddTask)

	http.ListenAndServe(":8000", router)
}
