package daemon

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ABHINAV-JHA-27/scheduleX/internal/config"
)

func StartServer() {
	fmt.Println("Daemon Server Starting on PORT: ", config.Cfg.Daemon.Port)
	err := http.ListenAndServe(":"+strconv.Itoa(config.Cfg.Daemon.Port), nil)
	if err != nil {
		log.Fatalf("error starting server on port: %d", config.Cfg.Daemon.Port)
	}
}
