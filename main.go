package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")
	router := NewRouter()

	http.ListenAndServe(":8000", router)
}

type Router struct {
	routes map[string]http.HandlerFunc
}

func (r *Router) AddRoute(pattern string, handler http.HandlerFunc) {
	r.routes[pattern] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
		
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]http.HandlerFunc),
	}
}
