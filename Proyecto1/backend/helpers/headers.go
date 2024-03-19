package helpers

import "net/http"

func SetToJSON(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
}

func Cors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
