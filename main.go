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
	router.AddRoute(http.MethodGet, th.List.String(), th.ListTasks)
	router.AddRoute(http.MethodPost, th.Add.String(), th.AddTask)
	router.AddRoute(http.MethodPost, th.Show.String(), th.ShowTask)
	router.AddRoute(http.MethodDelete, th.Del.String(), th.DeleteTask)
	router.AddRoute(http.MethodPut, th.Edit.String(), th.EditTask)
	router.AddRoute("GET", "/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/swagger.json")
	})

	// service info
	router.RoutesList()

	http.ListenAndServe(":8000", router)
}
