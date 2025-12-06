package service

import (
	"context"
	"log"
	"time"

	"github.com/farizziezhi/layered/internal/domain"
	"github.com/farizziezhi/layered/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, user domain.User) (domain.User, error)
	Login(ctx context.Context, user domain.User) (string, error)
}

type UserServiceImpl struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &UserServiceImpl{Repo: repo}
}

func (service *UserServiceImpl) Register(ctx context.Context, user domain.User) (domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("ERROR:", err)
		return user, err
	}
	user.Password = string(hashedPassword)
	return service.Repo.Register(ctx, user)
}

func (service *UserServiceImpl) Login(ctx context.Context, user domain.User) (string, error) {
	claims, _, err := service.Repo.Login(ctx, user)
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(30 * time.Minute)
	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(expirationTime)

	var jwtKey = []byte("alamak_kasus_ni") 
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
