package processhandler

import (
	processcontroller "backend/controllers/process.controller"
	"backend/helpers"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, "process/getAll")
	helpers.Cors(&w)
	helpers.SetToJSON(&w)
	switch r.Method {
	case http.MethodGet:
		processes, err := processcontroller.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		processesJSON, err := json.Marshal(processes)
		w.WriteHeader(http.StatusOK)
		w.Write(processesJSON)
	default:
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}

func New(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, "process/new")
	helpers.Cors(&w)
	helpers.SetToJSON(&w)
	switch r.Method {
	case http.MethodGet:
		processInfo, err := processcontroller.New()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		processJSON, err := json.Marshal(processInfo)
		w.WriteHeader(http.StatusOK)
		w.Write(processJSON)
	default:
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}

func Ready(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, "process/ready")
	helpers.Cors(&w)
	helpers.SetToJSON(&w)
	switch r.Method {
	case http.MethodGet:
		pidStr := r.URL.Query().Get("pid")
		pid, err := strconv.ParseInt(pidStr, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		processInfo, err := processcontroller.Ready(pid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		processJSON, err := json.Marshal(processInfo)
		w.WriteHeader(http.StatusOK)
		w.Write(processJSON)
	default:
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}

func Running(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, "process/running")
	helpers.Cors(&w)
	helpers.SetToJSON(&w)
	switch r.Method {
	case http.MethodGet:
		pidStr := r.URL.Query().Get("pid")
		pid, err := strconv.ParseInt(pidStr, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		processInfo, err := processcontroller.Running(pid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		processJSON, err := json.Marshal(processInfo)
		w.WriteHeader(http.StatusOK)
		w.Write(processJSON)
	default:
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}

func Waiting(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, "process/waiting")
	helpers.Cors(&w)
	helpers.SetToJSON(&w)
	switch r.Method {
	case http.MethodGet:
		pidStr := r.URL.Query().Get("pid")
		pid, err := strconv.ParseInt(pidStr, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		processInfo, err := processcontroller.Waiting(pid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		processJSON, err := json.Marshal(processInfo)
		w.WriteHeader(http.StatusOK)
		w.Write(processJSON)
	default:
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}

func Terminated(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, "process/terminated")
	helpers.Cors(&w)
	helpers.SetToJSON(&w)
	switch r.Method {
	case http.MethodGet:
		pidStr := r.URL.Query().Get("pid")
		pid, err := strconv.ParseInt(pidStr, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		processInfo, err := processcontroller.Terminated(pid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		processJSON, err := json.Marshal(processInfo)
		w.WriteHeader(http.StatusOK)
		w.Write(processJSON)
	default:
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}
