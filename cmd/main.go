package main

import (
	"log"
	"os"

	"github.com/devstackq/go-clean/config"
	"github.com/devstackq/go-clean/server"
	"github.com/spf13/viper"
	// _ "github.com/lib/pq"
)

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
	//refactor
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Print("logger start")
}
