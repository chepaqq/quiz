package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/chepaqq99/quiz/pkg/handler"
	"github.com/chepaqq99/quiz/pkg/server"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("config/")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error initializing configs: %s", err.Error())
	}
	srv := new(server.Server)
	fmt.Printf("Start server on %s port\n", viper.GetString("port"))
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalln(err.Error())
	}
}
