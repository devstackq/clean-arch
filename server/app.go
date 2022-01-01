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
	"github.com/devstackq/go-clean/delivery"

	handler "github.com/devstackq/go-clean/auth/delivery/http"
	mongoRepo "github.com/devstackq/go-clean/auth/repository/mongo"
	"github.com/devstackq/go-clean/auth/usecase"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type App struct {
	httpServer  http.Server
	grpcServer  grpc.Server
	authUseCase auth.UseCase
}

// interface {Signup, Signin}; stuct Grpc - own realize; struct http - own realize, grpcServer
// 1 create initServer - http & grpc; 2 add delivery proto, hanlder, etc

// TODO : flex transport layer - with fabric
// TODO docker compose - mongo/sql
// https://dev.to/itscosmas/how-to-set-up-a-local-development-workflow-with-docker-for-your-go-apps-with-mongodb-and-mongo-express-f99
// //return preapred repo-service
// https://github.com/bxcodec/go-clean-arch-grpc/blob/master/main.go

func NewApp() *App {
	//psql case
	// storage2 := db.NewPostgresStorage("postgres", "password", "localhost:", "5432", "testdb")
	// dbSql, err := storage2.InitPostgresDb()
	//log.Println(err,1)
	// repoSql := psql.NewUserRepository(dbSql)
	// log.Print(repoSql, "init psql", err)

	//mongo case
	storage := db.NewMongoStorage("mongo", "", "mongodb://mongodb:", "27017", "testdb")
	dbMongo, err := storage.InitMongoDb()
	log.Print(err, 1)
	repoMongo := mongoRepo.NewUserRepository(dbMongo, viper.GetString("mongo.user_collection"))
	log.Print(repoMongo, "mongo init")

	http := delivery.NewHttpServer()
	server := http.InitHttp("8000")

	return &App{
		authUseCase: usecase.NewAuthUseCase(repoMongo, []byte(viper.GetString("auth.hash_salt")), []byte(viper.GetString("auth.secret_key")), viper.GetDuration("auth.token_ttl")),
		httpServer:  server,
		// grpcServer: grpc,
	}
}
func (a *App) initRoutes() {
	//http || grpc
	hr := handler.NewHandler(a.authUseCase)
	http.HandleFunc("/signup", hr.SignUp) //register handler
}

func (app *App) Run(port string) error {

	app.initRoutes()
	//app.InitGrpcRoutes()
	//grpc || http run
	go func() {
		if err := app.httpServer.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// //gracefull shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return app.httpServer.Shutdown(ctx)
}
