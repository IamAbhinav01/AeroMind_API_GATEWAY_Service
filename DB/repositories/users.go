package DB

import (
	"AeromindGO/models"
	"database/sql"
	"fmt"
	"log"
)

// "Any type that has a Create() method is a UserRepository."
type UserRepository interface {
	Create() (int,error)
	GetUserByID(id int) (models.User, error)
	GetAllUsers() ([]models.User, error)
	DeleteUserByID(id int) (sql.Result, error)
}

// the interface will be implemented by a struct

type UserRepositoryImpl struct {
	db *sql.DB // holds the memory address of thiss type
}

// Since UserRepositoryImpl has a Create() method, it implements the interface.
func (user *UserRepositoryImpl) Create()(int,error){
	
	query := "INSERT INTO users (id,email,password) VALUES(?,?,?)"
	output,err:=user.db.Exec(query,"2","abhinavSdd@xx.com","thisePasword")

	if err != nil{
		log.Fatal("Error inserting the user : ",err)
	}

	response,rowErr:=output.RowsAffected()

	if(rowErr != nil){
		log.Fatal("Error occured while checking rows affected . ",rowErr)
	}

	if response == 0{
		fmt.Println("No rows were affected , Error while creating the user")
	}
	return int(response), nil
}


func (user *UserRepositoryImpl) GetUserByID(id int) (models.User,error){
	
	var UsersModel models.User
	query:="SELECT id,email,password FROM USERS WHERE id = ?"
	// Query for a single row
	row:=user.db.QueryRow(query,id)

	err:=row.Scan(&UsersModel.Id,&UsersModel.Email,&UsersModel.Password)
	if err != nil{
		log.Fatal("Error while fetching the user by id : ",err)
	}

	return UsersModel,nil
}

func (user *UserRepositoryImpl) GetAllUsers() ([]models.User,error){

	query := "SELECT * from users"

	rows,err := user.db.Query(query)

	if err != nil{
		log.Fatal("Error while fetching all users : ",err)
	}
	
	defer rows.Close()

	var users []models.User

	for rows.Next(){
		var u models.User

		err := rows.Scan(&u.Id,&u.Email,&u.Password)
		if err != nil{
			log.Fatal("Error while scanning the user details : ",err)
		}
		users = append(users, u)
	}

	return users, nil
}

func (user *UserRepositoryImpl) DeleteUserByID(id int)(sql.Result,error){

	query := "DELETE FROM USERS WHERE id = ?"
	response,err:=user.db.Exec(query,id)
	if err != nil{
		log.Fatal("Error while deleting the user : ",err)
	}
	fmt.Println("User deleted successfully ")
	return response,nil

}

func NewUserRepository(_db *sql.DB) UserRepository{
	return &UserRepositoryImpl{
		db:_db,
	}
}