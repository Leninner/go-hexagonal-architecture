package main

import (
	"log"
	"net/http"

	"github.com/leninner/go-feature-flags/database"
	router "github.com/leninner/go-feature-flags/pkg"
	auth "github.com/leninner/go-feature-flags/pkg/auth/domain"
	authorsStore "github.com/leninner/go-feature-flags/pkg/users/data/store"
	authors "github.com/leninner/go-feature-flags/pkg/users/domain"
)

func main() {
	database.GetConnection()

	authorsRepo := authorsStore.New(database.DB)
	authorsSvc := authors.NewService(authorsRepo)
	authSvc := auth.NewService(authorsRepo)

	httpRouter := router.NewHTTPHandler(authorsSvc, authSvc)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", httpRouter), "Server failed")
}
