package main

import (
	"AeromindGO/app"
	DBconfig "AeromindGO/config/db"
	config "AeromindGO/config/env"
	"log"
)


func main(){

	config.Load()
	DBconfig.DBInit()
	app:=app.NewApplication()
	err:=app.Run()
	if(err != nil){
		log.Fatal(err)
	}
	app.Run()
}