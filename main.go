package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"streamer/server"
	"streamer/server/handler"
)

func main() {
	logger := log.New()
	logger.Println("Application Started")
	s := server.New(logger)
	router := handler.Router{
		Log: logger,
	}
	s.Register("/stream", router.Stream)
	s.Start(":8080")

	logger.Println("Application Shutdown")

	defer s.Shutdown(context.Background())
}
