package middlewares

import (
	"log"
	"net/http"

	"github.com/adailsonm/desafio-sword/lib"
	"github.com/gin-gonic/gin"
)

type DatabaseTrx struct {
	handler lib.RequestHandler
	db      lib.Database
}

func statusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}

func NewDatabaseTrx(
	handler lib.RequestHandler,
	db lib.Database,
) DatabaseTrx {
	return DatabaseTrx{
		handler: handler,
		db:      db,
	}
}

func (m DatabaseTrx) Setup() {
	log.Print("setting up database transaction middleware")

	m.handler.Gin.Use(func(c *gin.Context) {
		txHandle := m.db.DB.Begin()
		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

		c.Set("db_trx", txHandle)
		c.Next()

		if statusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated, http.StatusNoContent}) {
			if err := txHandle.Commit().Error; err != nil {
				log.Fatalf("trx commit error: ", err)
			}
		} else {
			txHandle.Rollback()
		}
	})
}
