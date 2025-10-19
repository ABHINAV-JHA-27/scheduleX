package main

import (
	"log"

	"github.com/ABHINAV-JHA-27/scheduleX/internal/cli"
	"github.com/ABHINAV-JHA-27/scheduleX/internal/config"
	"github.com/ABHINAV-JHA-27/scheduleX/internal/db"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %s", err)
	}
	err = db.InitDB()
	if err != nil {
		log.Fatalf("error initialising database: %s", err)
	}
	cli.Execute()
}
