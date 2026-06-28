package main

import (
	"AeromindGO/app"
	config "AeromindGO/config/env"
	"log"
)


func main(){


	
	config.Load()
	
	app:=app.NewApplication()
	err:=app.Run()
	if(err != nil){
		log.Fatal(err)
	}
	app.Run()
}