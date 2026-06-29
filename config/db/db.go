package config

import (
	config "AeromindGO/config/env"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func DBInit() (*sql.DB, error) {

	cfg := mysql.NewConfig()

	cfg.User = config.GetString("DBUSER","DBUSER")
	cfg.Passwd = config.GetString("DBPASS","DBPASS")
	cfg.Net = config.GetString("DB_Net","DB_Net")
	cfg.Addr = config.GetString("DB_Addr","DB_Addr")
	cfg.DBName = config.GetString("DBName","DBName")

	fmt.Println("Connecting to database : ",cfg.DBName,cfg.FormatDSN())

	db, err := sql.Open("mysql", cfg.FormatDSN())
	
	if err != nil{
		log.Fatal("Error occured while connecting the database. ",err)
	}

	pingErr := db.Ping()
	if pingErr != nil{
		log.Fatal("Error connecting to the database . ",err)
	}
	fmt.Println("Connected successfully to database .")

	return db,nil
}