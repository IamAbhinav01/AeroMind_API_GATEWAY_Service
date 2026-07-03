package middleware

import (
	"AeromindGO/dto"
	utils "AeromindGO/utils/json"
	formatters "AeromindGO/utils/responseFormatters"
	validators "AeromindGO/utils/validators"
	"net/http"
)

func LoginUserRequestValidation (next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.LoginUserDTO

		if err:= utils.FromJSON(r,&payload);err!=nil{
		formatters.ErrorResponse(w,http.StatusBadRequest,"Error occured while reading json.",err)
		return
		}

		if validationErr := validators.Validate.Struct(payload);validationErr != nil{
		formatters.ErrorResponse(w,http.StatusBadRequest,"Invalid request payload",validationErr)
		return
		}

		next.ServeHTTP(w,r)
	})
}

func CreateUserRequestValidation (next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.LoginUserDTO

		if err:= utils.FromJSON(r,&payload);err!=nil{
		formatters.ErrorResponse(w,http.StatusBadRequest,"Error occured while reading json.",err)
		return
		}

		if validationErr := validators.Validate.Struct(payload);validationErr != nil{
		formatters.ErrorResponse(w,http.StatusBadRequest,"Invalid request payload",validationErr)
		return
		}

		next.ServeHTTP(w,r)
	})
}