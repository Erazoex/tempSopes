package ramhandler

import (
	ramcontroller "backend/controllers/ram.controller"
	"backend/helpers"
	"encoding/json"
	"fmt"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, "ram/get")
	helpers.Cors(&w)
	helpers.SetToJSON(&w)
	switch r.Method {
	case http.MethodGet:
		ramInfo, err := ramcontroller.Get()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ramInfo.Total = (ramInfo.Total / 1024) / 1024
		ramInfo.Used = (ramInfo.Used / 1024) / 1024
		ramInfo.Free = (ramInfo.Free / 1024) / 1024
		ramJSON, err := json.Marshal(ramInfo)
		w.WriteHeader(http.StatusOK)
		w.Write(ramJSON)
	default:
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, "ram/getAll")
	helpers.Cors(&w)
	helpers.SetToJSON(&w)
	switch r.Method {
	case http.MethodGet:
		ramInfo, err := ramcontroller.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ramJSON, err := json.Marshal(ramInfo)
		w.WriteHeader(http.StatusOK)
		w.Write(ramJSON)
	default:
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}
