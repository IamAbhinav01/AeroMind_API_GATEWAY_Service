package DB

import (
	"database/sql"
	"fmt"
	"log"
)
// "Any type that has a Create() method is a UserRepository."
type UserRepository interface {
	Create()
}

// the interface will be implemented by a struct

type UserRepositoryImpl struct {
	db *sql.DB // holds the memory address of thiss type
}

// Since UserRepositoryImpl has a Create() method, it implements the interface.
func (user *UserRepositoryImpl) Create(){
	query := "INSERT INTO users (id,email,password) VALUES(?,?,?)"
	output,err:=user.db.Exec(query,"1","abhinav@xx.com","thisisthePasword")

	if err != nil{
		log.Fatal("Error inserting the user : ",err)
	}

	reponse,rowErr:=output.RowsAffected()

	if(rowErr != nil){
		log.Fatal("Error occured while checking rows affected . ",rowErr)
	}

	if reponse == 0{
		fmt.Println("No rows were affected , Error while creating the user")
	}

	fmt.Println("User created Successfully , rows affected : ",reponse)
}

func NewUserRepository(_db *sql.DB) UserRepository{
	return &UserRepositoryImpl{
		db:_db,
	}
}