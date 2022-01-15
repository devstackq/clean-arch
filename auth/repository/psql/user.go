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
	sqlQuery := `insert into users (username, password)values($1, $2) RETURNING id`
	row := ur.db.QueryRowContext(ctx, sqlQuery, user.Username, user.Password)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return err
	}
	log.Print("user create repo, pspl", id)
	return nil
}

func (ur UserRepository) GetUser(ctx context.Context, username, password string) (*models.User, error) {
	var user models.User
	sqlQuery := `select username, password from users where username = $1 and password = $2`
	row := ur.db.QueryRowContext(ctx, sqlQuery, username, password)
	err := row.Scan(&user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil

}
