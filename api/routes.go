package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var ic = NewIndexController()
var tc = NewTodosController(NewTodoRepo(getDBConnection()))

var routes = Routes{
	Route{
		"IndexController#Welcome",
		"GET",
		"/",
		ic.Welcome,
	},
	Route{
		"TodosController#Index",
		"GET",
		"/todos",
		tc.TodoIndex,
	},
	Route{
		"TodosController#Show",
		"GET",
		"/todos/{todoId}",
		tc.TodoShow,
	},
	Route{
		"TodosController#Create",
		"POST",
		"/todos",
		tc.TodoCreate,
	},
	Route{
		"TodosController#Update",
		"PUT",
		"/todos/{todoId}",
		tc.TodoUpdate,
	},
	Route{
		"TodosController#Delete",
		"DELETE",
		"/todos/{todoId}",
		tc.TodoDelete,
	},
}
