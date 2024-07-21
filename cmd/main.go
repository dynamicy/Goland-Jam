package main

import (
	"Goland-Jam/pkg/config"
	"Goland-Jam/pkg/routes"
	"context"
	"log"
	"net/http"

	_ "Goland-Jam/docs" // 引入 Swagger docs
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Goland-Jam API
// @version 1.0
// @description This is a sample server for Goland-Jam.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

func main() {
	cfg := config.LoadConfig()
	client := config.ConnectDB(cfg.MongoURI)
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect MongoDB client: %v", err)
		}
	}()

	routes.SetupRoutes(client)

	// 設置 Swagger 路由
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	log.Printf("Starting server on :%s\n", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
