package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/farizziezhi/layered/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Register(ctx context.Context, user domain.User) (domain.User, error)
	Login(ctx context.Context, user domain.User) (domain.Claims, domain.User, error)
}

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}


func (repo *UserRepositoryImpl) Register(ctx context.Context, user domain.User) (domain.User, error) {
	var cekUser domain.User
	err := repo.DB.QueryRowContext(ctx, "SELECT id, username, password FROM users WHERE username = ?", user.Username).Scan(&cekUser.ID, &cekUser.Username, &cekUser.Password)
	if err == nil {
		log.Println("ERROR: username sudah ada")
		return user, errors.New("username sudah ada")
	}

	result, err := repo.DB.ExecContext(ctx, "INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		log.Println("ERROR:", err)
		return user, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("ERROR:", err)
		return user, err
	}
	user.ID = int(id)

	return user, nil
}

func (repo *UserRepositoryImpl) Login(ctx context.Context, user domain.User) (domain.Claims, domain.User, error) {
	var result domain.User
	err := repo.DB.QueryRowContext(ctx, "SELECT id, username, password FROM users WHERE username = ?", user.Username).Scan(&result.ID, &result.Username, &result.Password)
	if err != nil {
		log.Println("ERROR:", err)
		return domain.Claims{}, domain.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if err != nil {
		log.Println("ERROR: password not match")
		return domain.Claims{}, domain.User{}, err
	}

	claims := domain.Claims{
		Username: result.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: result.Username,
		},
	}

	return claims, result, nil
}

