package main

import (
	"healthcare-capt-america/pkg/configs"
	"healthcare-capt-america/pkg/databases"
	"log"
)

func main() {
	configPath, err := configs.ParseFlags()
	if err != nil {
		log.Fatal(err)
		return
	}
	config, err := configs.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	repo, err := databases.NewRepositories(config)
	if err != nil {
		return
	}
	databases.Automigration(repo)
}
