package main

import (
	"net/http"

	"github.com/Chethu16/todo-new-project/database"
	"github.com/Chethu16/todo-new-project/routers"
	"github.com/gorilla/mux"
)

func main(){
	databaseconnection:=database.DatabaseConnection("postgresql://car_parking_24cn_user:9tockjLfjrGkdafw01ckDZbakdXUyPDM@dpg-d1q8n849c44c739b0120-a.oregon-postgres.render.com/car_parking_24cn")
	defer databaseconnection.Close()
	database.Initializing(databaseconnection)

	var route=mux.NewRouter()
	routers.InitializeRoutes(route,databaseconnection)
	



	http.ListenAndServe(":8000",route)
}