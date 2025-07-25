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


func(rp *Repo) Register(w http.ResponseWriter, r *http.Request){
	var userDetails models.User
    err:= json.NewDecoder(r.Body).Decode(&userDetails)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"message":"Json Dcode Error"})

	}
	enp,err:=bcrypt.GenerateFromPassword([]byte(userDetails.UserPassword),bcrypt.DefaultCost)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"message":"Password Encrypt Error"})
		return
		
	}
	
	_,err=rp.Db.Exec(`INSERT INTO users VALUES ($1,$2,$3,$4)`,uuid.NewString(),userDetails.UserName,userDetails.UserEmail,string(enp))
		if err!=nil{
			json.NewEncoder(w).Encode(map[string]string{"message":"Query execution error"})
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"message":"Registerd Succesfully"})
}
func (rp *Repo)Login(w http.ResponseWriter,r *http.Request){
	var login models.User
	err:=json.NewDecoder(r.Body).Decode(&login)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"message":"json login decode error"})
		return
	}
	var dbpassword,dbuserid string
	err=rp.Db.QueryRow(`SELECT user_id,user_password FROM users WHERE user_email=$1`,login.UserEmail).Scan(&dbuserid,&dbpassword)
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]string{"message":"password scan error"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbpassword) , []byte(login.UserPassword))
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"message":"Password incorrect"})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"user_id":dbuserid})
}