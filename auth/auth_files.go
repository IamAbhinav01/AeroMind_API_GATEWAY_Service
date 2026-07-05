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

	jwt_token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt_payload)

	Token, Tokenerr:=jwt_token.SignedString(JWT_SECRET)

	if Tokenerr != nil{
		fmt.Println("Error signing token:", Tokenerr)
		return "", Tokenerr
	}

	return Token,nil

}


func Verify_JWT_TOKEN(token string) (*jwt.Token,error){

	jwt_token,err:= jwt.Parse(token,func(t *jwt.Token) (any, error) {
		return JWT_SECRET,nil
	})

	if err != nil{
		fmt.Println("Invalid token")
		return nil,err
	}

	if !jwt_token.Valid {
		fmt.Println("Invalid token")
		return nil,err
	}

	return jwt_token,nil

}