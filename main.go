package main

import (
	"net/http"

	h "github.com/ovguschin90/todolist/internal/handler"
	th "github.com/ovguschin90/todolist/internal/handler/todo"
	"github.com/ovguschin90/todolist/internal/router"
)

func main() {
	router := router.New()

	//AddRoutes
	router.AddRoute(http.MethodGet, "/", h.Index)
	router.AddRoute(http.MethodGet, "/todos", th.List)
	// router.AddRoute(http.MethodGet, "/todos/id=[0-9]+", th.ShowTask)
	// router.AddRoute(http.MethodPost, "/todos/add", th.AddTask)
	// router.AddRoute(http.MethodPost, "/todos/del", th.DeleteTask)

	router.RoutesList()

	http.ListenAndServe(":8000", router)
}
