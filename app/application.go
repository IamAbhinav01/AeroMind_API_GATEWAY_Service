package app

import (
	config "AeromindGO/config/env"
	"fmt"
	"net/http"
	"time"
)

type Config struct{
	Addr string
}

type Application struct{
	Config Config
}

func NewApplication() *Application{
	return  &Application{
		Config: Config{
			Addr: config.GetString("PORT",":3000"),
		},
	}
}

func (app *Application) Run() error{
	server:=&http.Server{
		Addr: app.Config.Addr,
		Handler: nil,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
	}
	fmt.Println("Starting server on ",app.Config.Addr)

	return server.ListenAndServe()
}
