package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/leninner/go-feature-flags/pkg/users/domain"
)

// NewRoutesFactory create and returns a factory to create routes for the users module
func NewRoutesFactory(group *gin.RouterGroup) func(service domain.UserService) {
	userRoutesFactory := func(service domain.UserService) {
		group.GET("/:id", func(c *gin.Context) {
			id := c.Query("id")
			results, err := service.GetUser(id)

			if err != nil {
				c.Error(err)
				return
			}

			c.JSON(http.StatusOK, results)
		})

	}

	return userRoutesFactory
}
