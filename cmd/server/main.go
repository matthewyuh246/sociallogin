package main

import (
	"github.com/matthewyuh246/socallogin/internal/router"
	db "github.com/matthewyuh246/socallogin/pkg/database"
)

func main() {
	db := db.NewDB()
	e := router.NewRouter()
	e.Logger.Fatal(e.Start(":8080"))
}
