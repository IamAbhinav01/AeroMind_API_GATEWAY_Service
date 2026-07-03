package auth

import (
	config "AeromindGO/config/env"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var JWT_SECRET = []byte(config.GetString("JWT_SECRET","JWT_SECRET"))

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),14)

	if err != nil{
		log.Fatal("Error occured while hashing password ",err)
	}

	return string(bytes),nil

}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Generate_JWT_TOKEN(email string,id int) (string,error){
	jwt_payload := jwt.MapClaims{
		"email": email,
		"id":id,
	}

	jwt_token := jwt.NewWithClaims(jwt.SigningMethodES256,jwt_payload)

	Token, Tokenerr:=jwt_token.SignedString(JWT_SECRET)

	if Tokenerr != nil{
		fmt.Println("Error signing token:", Tokenerr)
		return "", Tokenerr
	}

	return Token,nil

}
