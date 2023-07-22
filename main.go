package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	pkg "github.com/seedy1/UpfTest/packages"
	"github.com/seedy1/UpfTest/routes"
)

func main() {
	config, err := LoadConfig("config.json")
	if err != nil {
		log.Fatal("main - Error loading config file", err)
		os.Exit(1)
	}
	log.Println("Starting server on port", config.Port)
	http.HandleFunc("/analysis", routes.HandleSSE)
	log.Fatal(http.ListenAndServe(config.Port, nil))

}

// load configuration file
func LoadConfig(file string) (pkg.Config, error) {
	var config pkg.Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		log.Fatal("loadConfig - Error loading config file", err)
		return pkg.Config{}, errors.New("loadConfig - Error loading config file")
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, nil
}
