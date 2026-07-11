package app

import (
	DB "AeromindGO/DB/repositories"
	DBconfig "AeromindGO/config/db"
	config "AeromindGO/config/env"
	"AeromindGO/controllers"
	"AeromindGO/middleware"
	"AeromindGO/router"
	"AeromindGO/services"
	"AeromindGO/utils/limiter"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
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

	var redisClient *redis.Client
	redisURL := config.GetString("REDIS_URL", "")
	if redisURL != "" {
		opt, err := redis.ParseURL(redisURL)
		if err != nil {
			fmt.Println("Error parsing REDIS_URL:", err)
			return err
		}
		redisClient = redis.NewClient(opt)
	} else {
		redisClient = redis.NewClient(&redis.Options{
			Addr: config.GetString("REDIS_ADDR", "localhost:6379"),
		})
	}


	ur:=DB.NewUserRepository(db)
	us:=services.NewUserService(ur)
	uc:=controllers.NewUserController(us)
	

	rateService:=limiter.NewRedisBucketService(redisClient)

	rateLimiter := middleware.CreateRateLimiter(
    rateService,
    middleware.RateLimiterOptions{
        Label:    "standard",
        Capacity: 5,
        Refill:   5,
        Timeout:  10000,
        Cost:     1,
        Whitelist: []string{"127.0.0.1"},
    },
	
	)
	
	uRouter:=router.NewRouter(uc,rateLimiter)

	server:=&http.Server{
		Addr: app.Config.Addr,
		Handler: router.SetupRouter(uRouter),
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
	}
	fmt.Println("Starting server on ",app.Config.Addr)

	return server.ListenAndServe()
}
