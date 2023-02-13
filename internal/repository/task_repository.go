package repository

import (
	"log"

	"github.com/adailsonm/desafio-sword/lib"
	"gorm.io/gorm"
)

type TaskRepository struct {
	lib.Database
}

func NewTaskRepository(db lib.Database) TaskRepository {
	return TaskRepository{
		Database: db,
	}
}

func (r TaskRepository) WithTrx(trxHandle *gorm.DB) TaskRepository {
	if trxHandle == nil {
		log.Fatal("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}
