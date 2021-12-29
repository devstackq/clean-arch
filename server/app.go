package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/devstackq/go-clean/auth"
	mongoRepo "github.com/devstackq/go-clean/auth/repository/mongo"
	"github.com/devstackq/go-clean/auth/usecase"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//config server; app; init db
//create custom server
type App struct {
	httpServer  http.Server
	authUseCase auth.UseCase
}
Server Run; docker - mongo start

//return preapred repo-service
func NewApp() *App {
	db := InitDb()
	userRepo := mongoRepo.NewUserRepository(db, viper.GetString("mongo.user_collection"))

	return &App{
		authUseCase: usecase.NewAuthUseCase(userRepo, []byte(viper.GetString("auth.hash_salt")), []byte(viper.GetString("auth.secret_key")), viper.GetDuration("auth.token_ttl")),
	}
}

func InitDb() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("mongo.uri")))
	if err != nil {
		log.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Print(err)
	}

	return client.Database(viper.GetString("mongo.name"))

}
