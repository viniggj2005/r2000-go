package main

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/viniggj2005/r2000-go/clients"
	// "github.com/viniggj2005/r2000-go/mqtt"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar env .env file: %s", err)
	}
	//##################################################################
	//CLIENTE MQTT
	// client := mqtt.ConnectMqtt()
	// mqtt.Sub(client, "teste")
	// mqtt.Publish(client, "topic/test", 0, true, []byte(`{"teste":"vini"}`))]
	//##################################################################
	//PARTE CONEXÃO SERIAL
	// client := serial.GetPorts()
	// fmt.Println("portas:", client)
	// port, err := serial.SerialConnection(client[0])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer port.Close()

	// _, _ = port.Write([]byte{0xA0, 0x03, 0x01, 0x72, 0xEA})
	// for {
	// 	buf := make([]byte, 128)
	// 	n, err := port.Read(buf)
	// 	if err != nil {
	// 		log.Fatal("Erro ao ler da porta:", err)
	// 	}
	// 	if n > 0 {
	// 		fmt.Printf("Recebido: % X\n", buf[:n])
	// 	}
	// }
	//##################################################################
	//CRIAR O FRAME APARTIR DE PARAMETROS
	// utils.BuildCommandFrame(utils.BuildFrame{Command: enums.GET_READER_TEMPERATURE, Params: []byte{byte(enums.ANTENNA1)}})
	//##################################################################

	callbacks := clients.OnReadingCallbacks{
		OnFirmware: func(client *clients.R2000Client, version float64) {
			fmt.Println("Firmware version:", version)
		},
		OnTemperature: func(client *clients.R2000Client, temperature int) {
			fmt.Println("Temperature:", temperature)
		},
		OnGetDrmStatus: func(client *clients.R2000Client, status string) {
			fmt.Println("DRM Status:", status)
		},
		// você adiciona os outros conforme necessário
	}

	client, err := clients.NewR2000Client("teste", "COM7", callbacks)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// loop infinito enviando comandos
	for {
		// aqui você chama uma função que escreve na serial
		// Exemplo (precisa existir no seu pacote connection/utils):
		// connection.SendCommand(client.Port, enums.GET_READER_TEMPERATURE)

		fmt.Println("Comando enviado, aguardando resposta...")
		time.Sleep(2 * time.Second) // espera um pouco antes de mandar de novo
	}

}
