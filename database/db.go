package database

import (
	"database/sql"
	"log"
	_"github.com/lib/pq"
)


func DatabaseConnection(url string) *sql.DB{
	cnn,err:=sql.Open("postgres",url)
	if err!=nil{
		log.Fatalf("unable to connect database :%v",err)
	}
	log.Println("Database connected Succesfully")
	return cnn
}

func Initializing(databaseconnection *sql.DB){
	var queries =[]string{
		`CREATE TABLE IF NOT EXISTS users(
		user_id VARCHAR NOT NULL,
		user_name VARCHA NOT NULL,
		user_email VARCHAR NOT NULL,
		user_password VARCHAR NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS todo(
		todo_title VARCHAR NOT NULL,
		todo_description VARCHAR NOT NULL,
		todo_status BOOLEAN NOT NULL
		)`,
	
	}
	for _,query:=range queries{
		_,err :=databaseconnection.Exec(query)
		if err!=nil{
			log.Fatalf("Unable to initialize :%v",err)

		}

	}
	log.Println("Database Initialized SUccesfully")
}