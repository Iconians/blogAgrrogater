package main

import (
	"fmt"
	"log"

	"gatorapp/internal/config"
)

func main() {
	// Read config
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	// Set current user
	err = cfg.SetUser("your_name") // <-- replace with your name
	if err != nil {
		log.Fatalf("failed to set user: %v", err)
	}

	// Read again
	updatedCfg, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	fmt.Println(updatedCfg.DBUrl)
	// Print config
	// fmt.Printf("Config: %+v\n", updatedCfg)
}
