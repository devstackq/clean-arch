package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/devstackq/go-clean/auth"
	authHttp "github.com/devstackq/go-clean/auth/delivery/http"
	mongoRepo "github.com/devstackq/go-clean/auth/repository/mongo"
	"github.com/devstackq/go-clean/auth/usecase"
	"github.com/devstackq/go-clean/db"
	"github.com/devstackq/go-clean/transport"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	// grpc        grpc.Server
	authUseCase auth.UseCase
}

// interface {Signup, Signin}; stuct Grpc - own realize; struct http - own realize, grpcServer
//singletone - prepare app, connect layers with interface; init app

func NewApp() *App {
	//psql case
	// storage2 := db.NewPostgresStorage("postgres", "password", "localhost:", "5432", "testdb") os.LookUp
	// dbSql, err := storage2.InitPostgresDb()
	// repoSql := psql.NewUserRepository(dbSql)

	//mongo case
	// storage := db.NewMongoStorage("mongo", "", "mongo", "27017", "testdb") // os.LookUp
	// dbMongo, err := storage.InitMongoDb()

	//2 variant db, method fabric
	mongoObject := db.NewDbObject("mongodb", viper.GetString("mongo.username"), viper.GetString("mongo.password"), viper.GetString("mongo.uri"), viper.GetString("mongo.port"), viper.GetString("mongo.dbName"), viper.GetString("mongo.user_collection"))
	// db.SetConfig("", "", viper.GetString("mongo.uri"), "27017", "users")
	db, err := mongoObject.InitDb()
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println("init db")

	repoMongo := mongoRepo.NewUserRepository(db.(*mongo.Database), viper.GetString("mongo.user_collection"))
	log.Print(repoMongo, "init repo")
	return &App{
		authUseCase: usecase.NewAuthUseCase(repoMongo, []byte(viper.GetString("auth.hash_salt")), []byte(viper.GetString("auth.secret_key")), viper.GetDuration("auth.token_ttl")),
		// httpServer:  server.(http.Server),
	}
}

func (app *App) Run(port string) error {
	//grpc || http create server
	factory := transport.GetFactory("http")
	transportProtocol := factory.GetTransport()
	server := transportProtocol.InitTransport(viper.GetString("port")).(http.Server)

	authHttp.InitRoutes(app.authUseCase)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	log.Print("run server port: ", viper.GetString("port"))

	//refactor logger go func()
	// file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND, 0644)
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer file.Close()
	// log.SetOutput(file)
	// log.Print("logger start")

	//gracefull shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return server.Shutdown(ctx)
}

//func NewServer(){}
