package main

import (
	"consumption_tracker/cmd/config"
	"consumption_tracker/cmd/internal/application/services"
	"consumption_tracker/cmd/internal/infrastructure/database/postgresql"
	"consumption_tracker/cmd/internal/infrastructure/httpclient"
	"consumption_tracker/cmd/internal/interfaces/http/handlers"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	if loadedConfig, err := config.LoadConfig(); err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		return
	} else {
		fmt.Printf("Configuration loaded: %v\n", loadedConfig)

		// Set up the database
		db, err := sql.Open("postgres", loadedConfig.DBURL)
		if err != nil {
			fmt.Printf("Error connecting to the database: %v\n", err)
			return
		}

		// Set up the repository
		repository := postgresql.NewPostgresqlRepository(db)

		httpClient := &http.Client{}
		addressService := httpclient.NewAddressClient(httpClient, loadedConfig.AddressServiceURL)

		// Initialize energy consumption service
		consumptionService := services.NewEnergyConsumptionService(repository, addressService)

		// Initialize handlers
		consumptionHandler := handlers.NewConsumptionHandler(*consumptionService)

		// Set up the router
		router := gin.Default()
		router.GET("/consumption", consumptionHandler.GetConsumption)

		// Start the server
		if err := router.Run(fmt.Sprintf(":%s", loadedConfig.ServerPort)); err != nil {
			fmt.Printf("Error starting server: %v\n", err)
		}
	}
}
