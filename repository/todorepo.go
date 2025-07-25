package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/Chethu16/todo-new-project/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type TodoStruct struct {
	DB *sql.DB
}

func (ts *TodoStruct) TodoRepo(w http.ResponseWriter, r *http.Request) {
	var todoDetails models.Todo
	err := json.NewDecoder(r.Body).Decode(&todoDetails)
	if err != nil {

		json.NewEncoder(w).Encode(map[string]string{"message": "Todo json decode error"})
		return
	}

	_, err = ts.DB.Exec(`INSERT INTO todo VALUES($1,$2,$3,$4,$5)`, todoDetails.UserID, uuid.NewString(), todoDetails.TodoTitle, todoDetails.TodoDescription, todoDetails.TodoStatus)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"message": "Query execution error"})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Todo added Successfully"})

}

func (ts *TodoStruct) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var UpdateDetails models.Todo
	err := json.NewDecoder(r.Body).Decode(&UpdateDetails)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"message": "Todo json update error"})
		return
	}

	_, err = ts.DB.Exec(`UPDATE todo SET todo_title=$1,todo_description=$2 WHERE todo_id=$3`, UpdateDetails.TodoTitle, UpdateDetails.TodoDescription, UpdateDetails.TodoID)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(map[string]string{"message": "Update query execution error"})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Updated Succesfully"})
}

func (ts *TodoStruct) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	var userId = vars["id"]

	_, err := ts.DB.Exec("UPDATE todo SET todo_status=$1 WHERE user_id=$2", true, userId)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"Message": "error occured"})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"Message": "Task Done"})
}

func (ts *TodoStruct) GetTodos(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	var userId = vars["id"]

	res, err := ts.DB.Query("SELECT todo_title , todo_description , todo_status , todo_id FROM todo WHERE user_id=$1", userId)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"Message": "error occured"})
		return
	}
	defer res.Close()
	var todos []models.Todo

	for res.Next() {
		var todo models.Todo
		err := res.Scan(&todo.TodoTitle, &todo.TodoDescription, &todo.TodoStatus, &todo.TodoID)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"Message": "error occured"})
			return
		}
		todos = append(todos, todo)
	}

	if res.Err() != nil {
		json.NewEncoder(w).Encode(map[string]string{"Message": "error occured"})
		return
	}

	json.NewEncoder(w).Encode(map[string][]models.Todo{"todos": todos})
}
func (ts *TodoStruct) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	var todoid = vars["todo_id"]
	_, err := ts.DB.Exec(`DELETE FROM todo WHERE todo_id=$1`, todoid)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"Message": "error occured"})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"Message": "Deleted Successfully"})
}
