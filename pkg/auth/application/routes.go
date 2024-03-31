package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/leninner/go-feature-flags/pkg/auth/domain"
	auth_dto "github.com/leninner/go-feature-flags/pkg/auth/dto"
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

			c.JSON(http.StatusOK, results)
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

			c.JSON(http.StatusOK, results)
		})

	}

	return userRoutesFactory
}
