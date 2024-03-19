package cpuhandler

import (
	cpucontroller "backend/controllers/cpu.controller"
	"backend/helpers"
	"encoding/json"
	"fmt"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, "cpu/get")
	helpers.Cors(&w)
	helpers.SetToJSON(&w)
	switch r.Method {
	case http.MethodGet:
		cpuInfo, err := cpucontroller.Get()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		cpuJSON, err := json.Marshal(cpuInfo)
		w.WriteHeader(http.StatusOK)
		w.Write(cpuJSON)
	default:
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, "cpu/getAll")
	helpers.Cors(&w)
	helpers.SetToJSON(&w)
	switch r.Method {
	case http.MethodGet:
		cpuInfo, err := cpucontroller.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		cpuJSON, err := json.Marshal(cpuInfo)
		w.WriteHeader(http.StatusOK)
		w.Write(cpuJSON)
	default:
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}
