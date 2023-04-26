package storage

import "todo/models"


type StorageI interface{
	Todo() TodoRepoI
}

type TodoRepoI interface{
	Create(entity models.CreateTodo)(err error)
}

