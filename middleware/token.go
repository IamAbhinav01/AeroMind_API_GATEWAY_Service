package middleware

import (
	"AeromindGO/auth"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader:= strings.Split(r.Header.Get("Authorization"), "Bearer ")

// headerPayload Fomat => 
	// authHeader = []string{
	//     "",
	//     "eyJhbGciOiJIUzI1NiIsInR5cCI...",
	// }


		if len(authHeader) != 2{
			fmt.Println("Malformed Header")
			http.Error(w,"Invalid Header recieved",http.StatusBadRequest)
			return
		}else{
			token:=authHeader[1]
			jwt_token,err:=auth.Verify_JWT_TOKEN(token)
			
			if err != nil{
				fmt.Println("Invalid Token recieved")
				http.Error(w,"Invalid Token recieved",http.StatusUnauthorized)
				return
			}

			claims,ok:=jwt_token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}


			userId,okId:=claims["id"].(float64)
			userEmail,okEmail:=claims["email"].(string)

			if !okId || !okEmail {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
			}

			fmt.Println("Authenticated user ID:", int64(userId), "Email:", userEmail)
			

			ctx:=context.WithValue(r.Context(),"userId",int64(userId))
			ctx =context.WithValue(ctx,"userEmail",userEmail)

			next.ServeHTTP(w,r.WithContext(ctx))
		}


		

	})
}