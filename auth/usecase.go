package auth

import (
	"context"

	"github.com/devstackq/go-clean/auth/models"
)

type UseCase interface {
	SignUp(ctx context.Context, username, password string) error
	SignIn(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) *models.User
	HashPassword(password string) string
}
