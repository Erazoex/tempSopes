package routes

import (
	cpuhandler "backend/handlers/cpu.handler"
	"fmt"
	"net/http"
)

func CpuRoutes() {
	baseUrl := "/cpu"
	http.HandleFunc(fmt.Sprintf("%s/get", baseUrl), cpuhandler.Get)
	http.HandleFunc(fmt.Sprintf("%s/getAll", baseUrl), cpuhandler.GetAll)
}
