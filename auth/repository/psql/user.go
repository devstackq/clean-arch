package psql

import (
	"context"
	"database/sql"
	"log"

	"github.com/devstackq/go-clean/auth/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	log.Print("user create repo, pspl")
	return nil
}

func (ur UserRepository) GetUser(ctx context.Context, username, password string) (*models.User, error) {
	log.Print("get User repo, pspl")
	//select * from user where username = $1 and password = $2
	return nil, nil
}
