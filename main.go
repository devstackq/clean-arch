package main

import (
	"log"

	"github.com/devstackq/go-clean/config"
	"github.com/devstackq/go-clean/server"
	"github.com/spf13/viper"
)

// connect doker mongo

func main() {
	if err := config.Init(); err != nil {
		log.Println(err, "viper")
		return
	}
	app := server.NewApp()
	if err := app.Run(viper.GetString("port")); err != nil {
		log.Println(err)
		return
	}

}
