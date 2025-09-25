package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/viniggj2005/r2000-go/serial"
	// "github.com/viniggj2005/r2000-go/mqtt"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar env .env file: %s", err)
	}
	// client := mqtt.ConnectMqtt()
	// mqtt.Sub(client, "teste")
	// mqtt.Publish(client, "topic/test", 0, true, []byte(`{"teste":"vini"}`))
	serial.PortsList()
	// select {}

}
