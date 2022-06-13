package main

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"

	"github.com/jlogicsoftware/starterkit/internal/app/server"
	"github.com/jlogicsoftware/starterkit/internal/pkg/config"
	"github.com/jlogicsoftware/starterkit/internal/pkg/logger"
)

func main() {
	cfg := config.Get()
	fmt.Println(*cfg)

	srv, err := server.New()
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to create server")
	}
	if err := server.Start(srv); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start server")
	}
}
