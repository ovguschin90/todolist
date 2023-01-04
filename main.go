package main

import (
	"net/http"

	h "github.com/ovguschin90/todolist/internal/handler"
	th "github.com/ovguschin90/todolist/internal/handler/todo"
	"github.com/ovguschin90/todolist/internal/router"

	httpSwagger "github.com/swaggo/http-swagger"
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
	router.AddRoute(http.MethodGet, "/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/docs.json"),
	))

	// service info
	router.RoutesList()

	http.ListenAndServe(":8000", router)
}
