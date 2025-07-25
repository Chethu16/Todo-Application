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
	var todo =repository.TodoStruct{
		DB: connection,
	}
	r.HandleFunc("/register",repo.Register).Methods("POST")
	r.HandleFunc("/login",repo.Login).Methods("POST")
	r.HandleFunc("/addtodo",todo.TodoRepo).Methods("POST")
	r.HandleFunc("/updatetodo",todo.UpdateTodo).Methods("POST")
	r.HandleFunc("/updatestatus/{id}", todo.UpdateStatus).Methods("GET")
	r.HandleFunc("/gettodo/{id}" , todo.GetTodos).Methods("GET")
	r.HandleFunc("/deletetodo/{todo_id}",todo.DeleteTodo).Methods("GET")
}