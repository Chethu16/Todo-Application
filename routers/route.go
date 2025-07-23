package routers

import (
	"database/sql"

	"github.com/Chethu16/todo-new-project/repository"
	"github.com/gorilla/mux"
)


func InitializeRoutes(r *mux.Router,connection *sql.DB){
	var repo = repository.Repo{
		Db: connection,
	}
	r.HandleFunc("/register",repo.Register).Methods("POST")
}