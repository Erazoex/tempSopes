package routes

import (
	ramhandler "backend/handlers/ram.handler"
	"fmt"
	"net/http"
)

func RamRoutes() {
	baseUrl := "/ram"
	http.HandleFunc(fmt.Sprintf("%s/get", baseUrl), ramhandler.Get)
	http.HandleFunc(fmt.Sprintf("%s/getAll", baseUrl), ramhandler.GetAll)
}
