package main

import (
	"fmt"
	"net/http"

	"github.com/pwsnet/golang-api/api"
	"github.com/pwsnet/golang-api/application"
	"github.com/pwsnet/golang-api/lib"
)

var (
	environment = "development"
	version     = "0.1"
	port        = 3000
)

// @title Golang API
// @version 0.1
// @description PWS Golang API Template for an API Service.
//
// @contact.name PWS Support
// @contact.email dev@pomaretta.com
//
// @BasePath /
// @schemes http
//
func main() {

	app := &application.Application{
		Version:     version,
		Environment: lib.Environment(environment),
	}
	api := api.New(app)

	http.ListenAndServe(fmt.Sprintf(":%d", port), api)
}
