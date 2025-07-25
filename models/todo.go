package models

type Todo struct{
	UserID string `json:"user_id"`
	TodoID string `json:"todo_id"`
	TodoTitle string `json:"todo_title"`
	TodoDescription string `json:"todo_description"`
	TodoStatus bool `json:"todo_status"`
}