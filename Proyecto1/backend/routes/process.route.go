package routes

import (
	processhandler "backend/handlers/process.handler"
	"fmt"
	"net/http"
)

func ProcessRoutes() {
	baseUrl := "/process"
	http.HandleFunc(fmt.Sprintf("%s/getAll", baseUrl), processhandler.GetAll)
	http.HandleFunc(fmt.Sprintf("%s/new", baseUrl), processhandler.New)
	http.HandleFunc(fmt.Sprintf("%s/ready", baseUrl), processhandler.Ready)
	http.HandleFunc(fmt.Sprintf("%s/running", baseUrl), processhandler.Running)
	http.HandleFunc(fmt.Sprintf("%s/waiting", baseUrl), processhandler.Waiting)
	http.HandleFunc(fmt.Sprintf("%s/terminated", baseUrl), processhandler.Terminated)
}
