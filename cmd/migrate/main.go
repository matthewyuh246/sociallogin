package main

import (
	"fmt"

	domain "github.com/matthewyuh246/socallogin/internal/domain"
	db "github.com/matthewyuh246/socallogin/pkg/database"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&domain.User{})
}
