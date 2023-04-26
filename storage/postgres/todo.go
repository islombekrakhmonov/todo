package postgres

import (
	"todo/models"
	"todo/storage"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)


type todoRepo struct{
	db *sqlx.DB
	todo storage.TodoRepoI
}

func (t todoRepo) Create(entity models.CreateTodo)(err error){
	insertQuery := `INSERT INTO todo(
		title, 
		description
	) VALUES (
		$1,
		$2
	)`
	_,err = t.db.Exec(insertQuery,
		entity.Title,
		entity.Description,
)
return err
}