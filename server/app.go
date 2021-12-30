package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/devstackq/go-clean/auth"
	handler "github.com/devstackq/go-clean/auth/deliviry/http"
	mongoRepo "github.com/devstackq/go-clean/auth/repository/mongo"
	"github.com/devstackq/go-clean/auth/usecase"
	"github.com/devstackq/go-clean/dbFabric"
	"github.com/spf13/viper"

	"go.mongodb.org/mongo-driver/mongo"
)

//config server; app; init db
//create custom server
type App struct {
	httpServer  http.Server
	authUseCase auth.UseCase
}

// Server Run; docker - mongo start

//return preapred repo-service
func NewApp() *App {
	//	storage2 := dbFabric.NewPostgresStorage("postgres", "password", "localhost:", "5432", "testdb")
	//	psqlDb, err := storage2.InitDb()
	//	userPsqlRepo := psqlRepo.NewUserRepository(psqlDb.(*sql.DB))

	//maybe use data from viper ?
	storage := dbFabric.NewMongoStorage("mongo", "", "mongodb://mongodb:", "27017", "testdb")
	mgDb, err := storage.InitDb()

	if err != nil {
		log.Print(err, 1)
		return nil
	}
	// log.Print(mgDb.(*mongo.Database), "mongo db")
	userMongoRepo := mongoRepo.NewUserRepository(mgDb.(*mongo.Database), viper.GetString("mongo.user_collection"))
	return &App{
		authUseCase: usecase.NewAuthUseCase(userMongoRepo, []byte(viper.GetString("auth.hash_salt")), []byte(viper.GetString("auth.secret_key")), viper.GetDuration("auth.token_ttl")),
	}
}

func (app *App) Run(port string) error {
	hr := handler.NewHandler(app.authUseCase)
	//init  & setup ahndlers
	http.HandleFunc("/signup", hr.SignUp) //register handler
	// http.HandleFunc("/signin", hr.Signin) //register handler

	//custom http server
	app.httpServer = http.Server{
		Addr: ":" + port,
		// Handler: ,
		ReadTimeout:    10 * time.Second, // each 10 sec read
		WriteTimeout:   10 * time.Second, //each 10sec write
		MaxHeaderBytes: 1 << 20,          // max 20 mg
	}
	go func() {
		if err := app.httpServer.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	//gracefull shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return app.httpServer.Shutdown(ctx)
}
