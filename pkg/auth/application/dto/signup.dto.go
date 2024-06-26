package auth

type SignUpDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Username string `json:"username" validate:"required"`
	RoleID   int    `json:"roleId" validate:"required"`
}
