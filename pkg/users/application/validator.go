package users

import (
	"github.com/gin-gonic/gin"
	users "github.com/leninner/go-feature-flags/pkg/users/domain/entities"
)

// AuthorValidator is a struct used to validate the JSON payload representing an author.
type AuthorValidator struct {
	Name  string `binding:"required" json:"name"`
	Email string `binding:"required" json:"email"`
}

// Bind validates the JSON payload and returns a Author instance corresponding to the payload.
func Bind(c *gin.Context) (*users.User, error) {
	var json AuthorValidator
	if err := c.ShouldBindJSON(&json); err != nil {
		return nil, err
	}

	user := &users.User{
		Username: json.Name,
		Email:    json.Email,
	}

	return user, nil
}
