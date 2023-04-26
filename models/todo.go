package models

import "time"

type Todo struct{
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
	Completed_at *time.Time `json:"created_at"`
	Updated_at *time.Time `json:"updated_at"`
}

type CreateTodo struct{
	Title string `json:"title"`
	Description string `json:"description"`
}

type UpdateTodo struct{
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
}