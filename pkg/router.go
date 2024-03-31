package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	coreErrors "github.com/leninner/go-feature-flags/pkg/core/errors"

	authRoutes "github.com/leninner/go-feature-flags/pkg/auth/application"
	healthRoutes "github.com/leninner/go-feature-flags/pkg/health"
	userRoutes "github.com/leninner/go-feature-flags/pkg/users/application"

	auth "github.com/leninner/go-feature-flags/pkg/auth/domain"
	users "github.com/leninner/go-feature-flags/pkg/users/domain"
)

func NewHTTPHandler(userSvc users.UserService, authSvc auth.AuthService) http.Handler {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	router.Use(coreErrors.Handler)

	healthGroup := router.Group("/health")
	healthRoutes.NewRoutesFactory()(healthGroup)

	api := router.Group("/api")
	//api.Use(authMiddleware)

	userGroup := api.Group("/user")
	userRoutes.NewRoutesFactory(userGroup)(userSvc)

	authGroup := api.Group("/auth")
	authRoutes.NewRoutesFactory(authGroup)(authSvc)

	return router
}
