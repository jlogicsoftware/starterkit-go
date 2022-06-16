package main

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"

	"github.com/jlogicsoftware/starterkit/internal/app/server"
	"github.com/jlogicsoftware/starterkit/internal/pkg/config"
	"github.com/jlogicsoftware/starterkit/internal/pkg/logger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
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
