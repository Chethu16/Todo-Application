package repository

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Chethu16/todo-new-project/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)


type Repo struct{
	Db *sql.DB

}


func(db *Repo) Register(w http.ResponseWriter, r *http.Request){
	var userDetails models.User
    err:= json.NewDecoder(r.Body).Decode(&userDetails)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"message":"Json Dcode Error"})

	}
	enp,err:=bcrypt.GenerateFromPassword([]byte(userDetails.UserPassword),bcrypt.DefaultCost)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"message":"Password Encrypt Error"})
		
	}
	
	_,err=db.Db.Exec(`INSERT INTO users VALUES ($1,$2,$3,$4)`,uuid.NewString(),userDetails.UserName,userDetails.UserEmail,string(enp))
		if err!=nil{
			json.NewEncoder(w).Encode(map[string]string{"message":"Query execution error"})
		}
		json.NewEncoder(w).Encode(map[string]string{"message":"Registerd Succesfully"})
}