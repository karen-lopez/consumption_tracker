package main

import (
	"ConsumptionTracker/cmd/config"
	"fmt"
)

func main() {
	if loadConfig, err := config.LoadConfig(); err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		return
	} else {
		fmt.Printf("Configuration loaded: %v\n", loadConfig)
		fmt.Println("Start application")
	}
}
