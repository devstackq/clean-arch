package usecase

import (
	"context"
	"log"
	"time"

	"github.com/devstackq/go-clean/auth"
	"github.com/devstackq/go-clean/models"
)

//repo;
//func constructor
//DI - each db - own realize; - condition - interface
//AuthUseCase struct - for relation between - layers; interface  - poly, DI;
type AuthUseCase struct {
	userRepo       auth.UserRepository
	hashSalt       string
	signinKey      []byte
	expireDuration time.Duration
}

func NewAuthUseCase(userRepo auth.UserRepository, hashSalt string, signinKey []byte, tokenTTLSecond time.Duration) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signinKey:      signinKey,
		expireDuration: time.Second * tokenTTLSecond,
	}
}

func (auth *AuthUseCase) SignUp(ctx context.Context, username, password string) error {
	pwd := shal.New()
	user := new(models.User)
	user.Username = username
	user.Password = password
	err := auth.userRepo.CreateUser(ctx, user)
	log.Print("call service auth, use case,  Signup", err)
	if err != nil {
		return err
	}
	return nil
}
func (auth *AuthUseCase) SignIn(ctx context.Context, username, password string) (string, error) {

}
func (auth *AuthUseCase) ParseToken(ctx context.Context, accessToken string) *models.User {

}
