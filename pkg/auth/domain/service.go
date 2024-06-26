package auth

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/leninner/go-feature-flags/database"
	auth_dto "github.com/leninner/go-feature-flags/pkg/auth/application/dto"
	user_models "github.com/leninner/go-feature-flags/pkg/users/data/models"
	users "github.com/leninner/go-feature-flags/pkg/users/domain"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(l auth_dto.LoginDTO) (string, error)

	SignUp(l auth_dto.SignUpDTO) (string, error)

	HashAndSalt(password []byte) (string, error)

	HasCorrectPassword(databasePwd string, incomingPwd []byte) bool
}

type Service struct {
	repository users.UserRepository
	jwtService JwtService
}

func (s *Service) Login(l auth_dto.LoginDTO) (string, error) {
	user, err := s.repository.GetByEmail(l.Email)
	if err != nil {
		return "", err
	}

	if hasCorrectPassword := s.HasCorrectPassword(user.Password, []byte(l.Password)); !hasCorrectPassword {
		return "", errors.New("invalid password")
	}

	return s.jwtService.Sign(CustomJwtClaims{
		Email:    user.Email,
		Username: user.Username,
		Exp:      time.Now().Add(time.Hour * 72).Unix(), // 3 days
		Sub:      strings.Join([]string{user.Email, "sub"}, "-"),
		Iat:      time.Now().Unix(),
	})
}

func (s *Service) SignUp(l auth_dto.SignUpDTO) (string, error) {
	hashedPassword, err := s.HashAndSalt([]byte(l.Password))
	if err != nil {
		return "", err
	}

	user := &user_models.User{
		Username: l.Username,
		Password: hashedPassword,
		Email:    l.Email,
		RoleID:   l.RoleID,
	}

	fmt.Println(user.RoleID)

	if hasBeenCreated := database.DB.Create(&user); hasBeenCreated.Error != nil {
		return "", hasBeenCreated.Error
	}

	return "User created successfully", nil
}

func (s *Service) HashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash), err
}

func (s *Service) HasCorrectPassword(databasePwd string, incomingPwd []byte) bool {
	byteHash := []byte(databasePwd)

	err := bcrypt.CompareHashAndPassword(byteHash, incomingPwd)
	return err == nil
}

func NewService(repository users.UserRepository, jwtService JwtService) *Service {
	return &Service{repository: repository, jwtService: jwtService}
}
