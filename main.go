package main

import (
	"github.com/joho/godotenv"
	"github.com/saipulmuiz/go-project-starter/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.SetReportCaller(true)

	app := config.Init()
	config.Catch(app.Start())
}
