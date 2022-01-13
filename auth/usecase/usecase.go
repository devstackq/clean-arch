package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"log"
	"time"

	"github.com/devstackq/go-clean/auth"
	"github.com/devstackq/go-clean/auth/models"
)

//repo;
//func constructor
//DI - each db - own realize; - condition - interface
//AuthUseCase struct - for relation between - layers; interface  - poly, DI;
type AuthUseCase struct {
	userRepo       auth.UserRepositoryInterface
	HashSalt       []byte
	signinKey      []byte
	expireDuration time.Duration
}

func NewAuthUseCase(userRepo auth.UserRepositoryInterface, hashSalt []byte, signinKey []byte, tokenTTLSecond time.Duration) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		HashSalt:       hashSalt,
		signinKey:      signinKey,
		expireDuration: time.Second * tokenTTLSecond,
	}
}

func (auth *AuthUseCase) SignUp(ctx context.Context, user *models.User) error {
	// auth.HashSalt = auth.generateSalt(16) //salt, then save Db
	user.Password = auth.hashPassword(user.Password) //update password - to hash + salt
	log.Print("call service auth, use case,  Signup", user)
	return auth.userRepo.CreateUser(ctx, user)
}

func (auth *AuthUseCase) SignIn(ctx context.Context, username, password string) (string, error) {
	// dbPassword, err := auth.userRepo.GetUserPassword(ctx, username)
	inputHashedPwd := auth.hashPassword(password)

	user, err := auth.userRepo.GetUser(ctx, username, inputHashedPwd)
	if err != nil {
		return "", err
	}
	log.Print(user, "next step parse toke")
	// auth.ParseToken()
	return "", nil
}

func (auth *AuthUseCase) ParseToken(ctx context.Context, accessToken string) *models.User {
	return nil
}

func (auth *AuthUseCase) hashPassword(password string) string {
	sha1Hasher := sha1.New()
	pwdBytes := []byte(password)
	//append hased password, with salt
	pwdBytes = append(pwdBytes, auth.HashSalt...)

	sha1Hasher.Write(pwdBytes) //write bytes - to hasher

	return fmt.Sprintf("%x", sha1Hasher.Sum(nil)) //hashed password
	// base64EncodingPasswordHash := base64.URLEncoding.EncodeToString(hashPasswordBytes)
}

// mongo/psql download; write service check application
