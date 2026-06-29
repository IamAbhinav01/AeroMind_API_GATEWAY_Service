package app

import (
	DB "AeromindGO/DB/repositories"
	DBconfig "AeromindGO/config/db"
	config "AeromindGO/config/env"
	"AeromindGO/controllers"
	"AeromindGO/router"
	"AeromindGO/services"
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
	db,err:=DBconfig.DBInit()
	if err != nil{
		fmt.Println("Error while setting up database",err)
		return err
	}
	ur:=DB.NewUserRepository(db)
	us:=services.NewUserService(ur)
	uc:=controllers.NewUserController(us)
	uRouter:=router.NewRouter(uc)
	server:=&http.Server{
		Addr: app.Config.Addr,
		Handler: router.SetupRouter(uRouter),
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
	}
	fmt.Println("Starting server on ",app.Config.Addr)

	return server.ListenAndServe()
}
