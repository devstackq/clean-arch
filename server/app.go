package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/devstackq/go-clean/auth"
	"github.com/devstackq/go-clean/db"
	"github.com/devstackq/go-clean/transport"
	"go.mongodb.org/mongo-driver/mongo"

	authHttp "github.com/devstackq/go-clean/auth/delivery/http"
	mongoRepo "github.com/devstackq/go-clean/auth/repository/mongo"
	"github.com/devstackq/go-clean/auth/usecase"
	"github.com/spf13/viper"
)

type App struct {
	httpServer http.Server
	// grpc        grpc.Server
	authUseCase auth.UseCase
}

// interface {Signup, Signin}; stuct Grpc - own realize; struct http - own realize, grpcServer
// 1 create initServer - http & grpc; 2 add delivery proto, hanlder, etc

//todo context all layer
// TODO : flex transport layer - with fabric
// TODO docker compose - mongo/sql
// https://dev.to/itscosmas/how-to-set-up-a-local-development-workflow-with-docker-for-your-go-apps-with-mongodb-and-mongo-express-f99
// //return preapred repo-service
// https://github.com/bxcodec/go-clean-arch-grpc/blob/master/main.go

func NewApp() *App {
	//psql case
	// storage2 := db.NewPostgresStorage("postgres", "password", "localhost:", "5432", "testdb") os.LookUp
	// dbSql, err := storage2.InitPostgresDb()
	// log.Println(err,1)
	// repoSql := psql.NewUserRepository(dbSql)
	// log.Print(repoSql, "init psql", err)

	//mongo case
	// storage := db.NewMongoStorage("mongo", "", "mongo", "27017", "testdb") // os.LookUp
	// dbMongo, err := storage.InitMongoDb()

	//2 variant, fabric
	f := db.GetDbFactory("mongodb")
	f.SetConfig("user", "password", "mongodb://mongo", "27017", "users")
	db, err := f.InitDb()

	log.Print(db, err)

	repoMongo := mongoRepo.NewUserRepository(db.(*mongo.Database), viper.GetString("mongo.user_collection"))
	log.Print(repoMongo, "mongo init")
	// db, err := getDb("psql")

	factory := transport.GetFactory("http")
	protocol := factory.GetProtocol()
	server := protocol.InitTransport("8080")

	return &App{
		authUseCase: usecase.NewAuthUseCase(repoMongo, []byte(viper.GetString("auth.hash_salt")), []byte(viper.GetString("auth.secret_key")), viper.GetDuration("auth.token_ttl")),
		httpServer:  server.(http.Server),
		// grpcServer: grpc,
	}
}

func (app *App) Run(port string) error {
	//grpc || http run
	authHttp.InitRoutes(app.authUseCase)
	//app.InitGrpcRoutes()
	go func() {
		if err := app.httpServer.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	//refactor
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Print("logger start")

	//gracefull shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return app.httpServer.Shutdown(ctx)
}
