package repository

import (
	"log"

	"github.com/adailsonm/desafio-sword/lib"
	"gorm.io/gorm"
)

type UserRepository struct {
	lib.Database
}

func NewUserRepository(db lib.Database) UserRepository {
	return UserRepository{
		Database: db,
	}
}

func (r UserRepository) WithTrx(trxHandle *gorm.DB) UserRepository {
	if trxHandle == nil {
		log.Fatal("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}
