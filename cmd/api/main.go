package main

import (
	"context"
	"log"

	"github.com/rpuglielli/person-api/config"
	"github.com/rpuglielli/person-api/internal/adapter/repository"
	"github.com/rpuglielli/person-api/internal/domain/person/usecase"
	"github.com/rpuglielli/person-api/internal/infrastructure/database"
	"github.com/rpuglielli/person-api/internal/infrastructure/http/handler"
	"github.com/rpuglielli/person-api/internal/infrastructure/http/router"
)

func init() {
	config.LoadEnv()
}

func main() {
	ctx := context.Background()

	mongoClient, err := database.NewMongoDBConnection(ctx, config.DB_URL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer mongoClient.Disconnect(ctx)

	db := mongoClient.Database(config.DB_NAME)

	personRepo := repository.NewMongoDBPersonRepository(db)
	personUseCase := usecase.NewPersonUseCase(personRepo)
	personHandler := handler.NewPersonHandler(personUseCase)

	r := router.SetupRouter(personHandler)

	if err := r.Run(":" + config.APP_PORT); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
