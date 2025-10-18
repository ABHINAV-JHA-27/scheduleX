package main

import (
	"log"

	"github.com/ABHINAV-JHA-27/scheduleX/internal/cli"
	"github.com/ABHINAV-JHA-27/scheduleX/internal/config"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %s", err)
	}
	cli.Execute()
}
