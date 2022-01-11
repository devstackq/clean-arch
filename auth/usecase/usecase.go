package usecase

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"log"
	"math/rand"
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
	//is Good || call handler layer ?
	// refactor, use hashsalt value from config

	auth.HashSalt = auth.GenerateSalt(16) //salt, then save Db
	user.Password = auth.hashPassword(user.Password)

	log.Print("call service auth, use case,  Signup", user)
	return auth.userRepo.CreateUser(ctx, user)
	// return nil
}

func (auth *AuthUseCase) SignIn(ctx context.Context, username, password string) (string, error) {
	//get pwd by username form Db; ==  salt + hash(current password ) if ok -> handler ParseTOken ?
	// auth.userRepo.GetUser(ctx, username, password)
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

	sha1Hasher.Write(pwdBytes)               //write bytes - to hasher
	hashPasswordBytes := sha1Hasher.Sum(nil) //hashed password
	//base64 convert
	var base64EncodingPasswordHash = base64.URLEncoding.EncodeToString(hashPasswordBytes)
	return base64EncodingPasswordHash
}

func (auth *AuthUseCase) GenerateSalt(saltSize int) []byte {
	// s := auth.HashSalt
	var salt = make([]byte, saltSize)
	_, err := rand.Read(salt[:])
	if err != nil {
		// panic(err) //recovery
		recover()
	}
	return salt
}

//signin use func
func (auth *AuthUseCase) ComparePassword(hashedPassword, inputPassword string) bool {
	//reverse proccess
	return hashedPassword == inputPassword
}

// mongo/psql download; write service check application
