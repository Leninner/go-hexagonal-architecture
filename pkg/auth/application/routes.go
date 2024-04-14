package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	auth_dto "github.com/leninner/go-feature-flags/pkg/auth/application/dto"
	domain "github.com/leninner/go-feature-flags/pkg/auth/domain"
	"github.com/leninner/go-feature-flags/pkg/core"
)

func NewRoutesFactory(group *gin.RouterGroup) func(service domain.AuthService) {
	userRoutesFactory := func(service domain.AuthService) {
		group.POST("/signup", func(c *gin.Context) {
			var signupDTO auth_dto.SignUpDTO
			if err := c.ShouldBindJSON(&signupDTO); err != nil {
				c.Error(err)
				return
			}

			results, err := service.SignUp(signupDTO)
			if err != nil {
				c.Error(err)
				return
			}

			c.JSON(http.StatusOK, core.NewResponseMessage("User created", results))
		})

		group.POST("/login", func(c *gin.Context) {
			var loginDTO auth_dto.LoginDTO
			if err := c.ShouldBindJSON(&loginDTO); err != nil {
				c.Error(err)
				return
			}

			results, err := service.Login(loginDTO)
			if err != nil {
				c.Error(err)
				return
			}

			c.SetCookie("access_token", results, 3600, "/", "localhost", false, true)
		})
	}

	return userRoutesFactory
}
