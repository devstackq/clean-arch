package mongo

import (
	"context"
	"log"

	"github.com/devstackq/go-clean/auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       primitive.ObjectID `bson:"id, omitempty"` // mongodb uuid Используя string, мы не привязываемся к конкретному хранилищу, и всегда можем конвертировать uuid или int в строку.
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}

//convert struct user - to mongo struct
func toMongoUser(u *models.User) *User {
	return &User{
		Username: u.Username,
		Password: u.Password,
	}
}

//convert mongo struct  - to model.User
func toModel(u *User) *models.User {
	return &models.User{
		ID:       u.ID.Hex(),
		Username: u.Username,
		Password: u.Password,
	}
}

type UserRepository struct {
	db *mongo.Collection
}

//get from mongo databse - collection
func NewUserRepository(db *mongo.Database, colection string) *UserRepository {
	return &UserRepository{
		db: db.Collection(colection),
	}
}

func (ur UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	model := toMongoUser(user)

	res, err := ur.db.InsertOne(ctx, model)
	if err != nil {
		log.Println(err)
		return err
	}
	//set generated user id from mongo & convert to Hex & set user.ID
	user.ID = res.InsertedID.(primitive.ObjectID).Hex()
	log.Print("success insert new user by Id", user.ID)

	return nil
}

func (ur UserRepository) GetUser(ctx context.Context, username, password string) (*models.User, error) {
	user := new(User) //mongoUser struct
	//get data from mongo, write to mongo struct -> convert to model.User

	//find 1 object by keys {username, password}
	err := ur.db.FindOne(ctx, bson.M{
		"username": username,
		"password": password,
	}).Decode(user) // finded object - set mongo User

	if err != nil {
		return nil, err
	}
	//retrun converted mongo -  default user struct
	return toModel(user), nil
}

func (ur UserRepository) GetUserPassword(ctx context.Context, username string) (string, error) {
	return "", nil
}

//create db; create colection; crud query
