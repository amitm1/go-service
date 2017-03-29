package main

import (
	"net/http"

	"github.com/amitm1/go-service/healthcheck"
	"github.com/amitm1/go-service/misc"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Getcache",
		"GET",
		"/getcache",
		GetCache,
	},
}

//  This adds all the default service routes that we will want.  They all have a pathPrefix of "/service"
var serviceroutes = Routes{
	Route{
		"Healthcheck",
		"GET",
		"/health",
		healthcheck.HealthCheckHandler,
	},
	Route{
		"Healthcheck",
		"GET",
		"/health/down",
		healthcheck.DownHandler,
	},
	Route{
		"Healthcheck",
		"GET",
		"/health/up",
		healthcheck.UpHandler,
	},
	Route{
		"Swagger",
		"GET",
		"/swagger",
		misc.SwaggerHandler,
	},
	Route{
		"Dependencies",
		"GET",
		"/dependencies",
		misc.DependenciesHandler,
	},
}
