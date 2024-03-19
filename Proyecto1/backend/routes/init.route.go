package routes

import (
	"backend/helpers"
	"net/http"
)

func Init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		helpers.Cors(&w)
	})
	// rutas relacionadas al cpu
	CpuRoutes()
	// rutas relacionadas a la memoria ram
	RamRoutes()
}
