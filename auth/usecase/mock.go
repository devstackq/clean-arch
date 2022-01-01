package usecase

import (
	"context"

	"github.com/devstackq/go-clean/auth/models"
	"github.com/stretchr/testify/mock"
)

type AuthUseCaseMock struct {
	mock.Mock
}

func (a AuthUseCaseMock) SignUp(ctx context.Context, username, password string) error {
	args := a.Mock.Called(username, password)
	return args.Error(0)
}

func (a AuthUseCaseMock) SignIn(ctx context.Context, username, password string) (string, error) {
	args := a.Mock.Called(username, password)
	return args[0].(string), args.Error(1)

}
func (a AuthUseCaseMock) ParseToken(ctx context.Context, accessToken string) *models.User {
	args := a.Mock.Called(accessToken)
	return args[0].(*models.User)
}
