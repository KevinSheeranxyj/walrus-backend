package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"kevinsheeran/walrus-backend/config"
	"kevinsheeran/walrus-backend/router"
	"log"
)

// @title Survey
// @version 1.0
// @description backend survey for walrus
// @securityDefinition.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	// Read the YAML file
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %s", err)
	}

	// Unmarshal the YAML data into the Config struct
	err = yaml.Unmarshal(data, &config.Config)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s", err)
	}
	router := router.Router()

	router.Run(config.Config.System.Host + ":" + config.Config.System.Port)

}
