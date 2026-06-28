package main

import (
	"AeromindGO/app"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)


func main(){


	

	
	PORT:=":" + os.Getenv("PORT")



	cfg:=app.Config{
		Addr: PORT,
	}

	app:=app.Application{
		Config: cfg,
	}

	err = app.Run()
	if err != nil{
		log.Fatal(err)
	}
	
}