package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"fizzbuzz",
		"GET",
		"/fizzbuzz",
		FizzbuzzHandler,
	},
	Route{
		"example",
		"GET",
		"/fizzbuzz/example",
		ExampleHandler,
	},
}
