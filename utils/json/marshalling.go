package utils

import (
	"encoding/json"
	"net/http"
)

func ToJSON(w http.ResponseWriter, status int, data any) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)

}

func FromJSON(r *http.Request,result any) error{
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // disallow unknown fields
	return decoder.Decode(result)
}