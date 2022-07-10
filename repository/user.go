package repository

import (
	"context"

	"github.com/harrycoa/apirest-websockets.git/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	Close() error
}

var implementation UserRepository

func SetRepository(repository UserRepository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func Close() error {
	return implementation.Close()
}

/*
************* Abstracciones *************
Handler - GetUserByIdPostgresql ...
es una clase concreta // rompe los principios SOLID

implementacion demasiada concreta y cuando se cambie de base de datos tendra problemas, es por ello que se usa
las abstracciónes (interfaces y sus implementaciones) eso es el patron repository
*/
