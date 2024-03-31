package main

import (
	"log"
	"net/http"

	"github.com/leninner/go-feature-flags/database"
	router "github.com/leninner/go-feature-flags/pkg"
	auth "github.com/leninner/go-feature-flags/pkg/auth/domain"
	userStore "github.com/leninner/go-feature-flags/pkg/users/data/store"
	users "github.com/leninner/go-feature-flags/pkg/users/domain"
)

func main() {
	database.GetConnection()

	userRepository := userStore.New(database.DB)
	jwtSvc := auth.NewJwtService()

	userSvc := users.NewService(userRepository)
	authSvc := auth.NewService(userRepository, jwtSvc)

	httpRouter := router.NewHTTPHandler(userSvc, authSvc)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", httpRouter), "Server failed")
}
