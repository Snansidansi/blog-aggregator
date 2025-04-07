package main

import (
	"fmt"
	"log"

	"github.com/snansidansi/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("snansidansi")
	if err != nil {
		log.Fatalf("error setting current user: %v\n", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}
	fmt.Printf("read config after changing current username: %+v\n", cfg)
}
