package utils

import (
	"encoding/json"
	"net/http"
)







func toJSON(w http.ResponseWriter,status int,data any)error{

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)

}





func SuccessResponse(w http.ResponseWriter,status int,message string,data any) error{
	response := map[string]any{}
	response["status"] = status
	response["message"] = message
	response["data"] = data
	return toJSON(w,status,response)
}
func ErrorResponse(w http.ResponseWriter,status int,message string,err error) error{
	response := map[string]any{}
	response["status"] = status
	response["message"] = message
	if err != nil {
		response["error"] = err.Error()
	} else {
		response["error"] = nil
	}
	return toJSON(w,status,response)
}