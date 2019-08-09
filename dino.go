package main

import (
	"dino/dinowebportal"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type configuration struct {
	ServerAddress string `json:"webserver"`
}

func main() {

	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}
	config := new(configuration)
	json.NewDecoder(file).Decode(config)
	fmt.Println(config.ServerAddress)
	dinowebportal.RunWebPortal(config.ServerAddress)

	//file
}
