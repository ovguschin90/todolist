package main

import (
	"net/http"
	"regexp"
)

func main() {
	router := &Router{}

	//AddRoutes
	router.AddRoute("^/$", index)

	http.ListenAndServe(":8000", router)
}

type Route struct {
	pattern *regexp.Regexp
	handler http.HandlerFunc
}

type Router struct {
	routes []*Route
}

func (rt *Router) AddRoute(pattern string, handler http.HandlerFunc) {
	route := &Route{
		pattern: regexp.MustCompile(pattern),
		handler: handler,
	}
	rt.routes = append(rt.routes, route)
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	for _, route := range rt.routes {
		if route.pattern.MatchString(path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}

	http.NotFound(w, r)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
