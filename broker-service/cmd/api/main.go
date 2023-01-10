package main

import (
	"fmt"
	"log"
)

const webPort = "80"

type Config struct{}

func main() {

	app := Config{}

	log.Printf("Starting broker service on port %s \n", webPort)

	// define http server

	srv := &http.Server{
		Addr:  fmt.Sprintf(": %s", webPort),
		Handler: app.routes()
	}

	// start the server
	err := srv.ListnerAndServe()
	if err != nil {
		log.Panic(err)
	}
}