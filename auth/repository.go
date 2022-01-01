package auth

import (
	"context"

	"github.com/devstackq/go-clean/auth/models"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, username, password string) (*models.User, error)
}

// определили основные сценарии работы системы регистрации/авторизации и описали абстракции для хранилища и бизнес логики
