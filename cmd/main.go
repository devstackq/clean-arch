package main

import (
	"net/http"

	authHttp "github.com/devstackq/go-clean/auth/deliviry/http"
	"github.com/devstackq/go-clean/auth/repository/mongo"
	"github.com/devstackq/go-clean/auth/usecase"
	"github.com/spf13/viper"
)

func main() {
	// ctx := context.Background()
	//signin service
	//d.HashSalt = d.getSaltByEmail(email)
	// hashPasswordDb := d.GetHasedPasswordByEmail(email)

	// inputHash := d.HashPassword(pwd)
	// state := d.ComparePassword(hashPasswordDb, inputHash)
	// if state {
	// 	d.SignIn(ctx, "madina", "123")
	// }

	repo := mongo.NewUserRepository(nil, viper.GetString("mongo.user_collection")) //set mongo db

	service := usecase.NewAuthUseCase(
		repo,
		[]byte(viper.GetString("auth.hash_salt")),
		[]byte(viper.GetString("auth.signin_key")),
		viper.GetDuration("auth.token_ttl"),
	)

	hr := authHttp.NewHandler(service)
	http.HandleFunc("/signup", hr.SignUp)
	http.ListenAndServe(":8000", nil)

}
