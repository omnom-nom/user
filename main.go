package main

import (
	"time"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/omnom-nom/user/api"
)

func main() {

	err := api.Init()
	if err != nil {
		log.Fatalf("Error %v Starting Server \n", err)
		time.Sleep(1 * time.Second)
		os.Exit(1)
	}

}
