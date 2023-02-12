package main

import (
	"github.com/adailsonm/desafio-sword/bootstrap"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	bootstrap.RootApp.Execute()
}
