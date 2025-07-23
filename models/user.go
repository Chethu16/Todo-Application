package models


type User struct{
	UserId string `json:"user_id"`
	UserName string `json:"user_name"`
	UserEmail string `json:"user_email"`
	UserPassword string `json:"user_password"`
}