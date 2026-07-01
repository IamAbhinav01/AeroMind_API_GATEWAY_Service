package utils

import "net/http"

func successResponse(w http.ResponseWriter,status int,message string,data any){
	response := map[string]any{
		response["status"] = status
		response["message"] = message
		response["data"] = data
		
	}
}
func errorResponse(w http.ResponseWriter,status int,message string,err error){
	response := map[string]any{
		response["status"] = status
		response["message"] = message
		response["error"] = err

	}
}