package main

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/viniggj2005/r2000-go/clients"
	"github.com/viniggj2005/r2000-go/dtos"
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
	// client := connection.GetPorts()
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

	callbacks := dtos.OnReadingCallbacks{
		OnFirmware: func(client dtos.R2000ClientIface, version float64) {
			fmt.Println("Firmware version:", version)
		},
		OnTemperature: func(client dtos.R2000ClientIface, temperature int) {
			fmt.Println("Temperature:", temperature)
		},
		OnGetDrmStatus: func(client dtos.R2000ClientIface, status string) {
			fmt.Println("DRM Status:", status)
		},
		OnGetFrequencyRegion: func(client dtos.R2000ClientIface, region string, frequency1, frequency2 float64, err error) {
			fmt.Println("region:", region, "f1:", frequency1, "f2:", frequency2)
		},
		OnGetOutputPower: func(client dtos.R2000ClientIface, power map[string]int) {
			for chave, valor := range power {
				fmt.Println("antenas:", chave, "potência :", valor, "Dbm")
			}
		},
		OnGetWorkAntenna: func(client dtos.R2000ClientIface, antenna string) {
			fmt.Println("region:", antenna)
		},
		OnSetBuzzerBehavior: func(client dtos.R2000ClientIface, ok bool, errMsg string) {
			fmt.Println("Comportamento setado:", ok)
		},
		OnSetFrequencyRegion: func(client dtos.R2000ClientIface, ok bool, errMsg string) {
			fmt.Println("Comportamento setado:", ok)
		},
		OnSetDrm: func(client dtos.R2000ClientIface, ok bool, errMsg string) {
			fmt.Println("Comportamento setado:", ok)
		},
		OnSetOutputPower: func(client dtos.R2000ClientIface, ok bool, errMsg string) {
			fmt.Println("Comportamento setado:", ok)
		},
		// OnSetWorkAntenna: func(client dtos.R2000ClientIface, ok bool, errMsg string) {
		// 	fmt.Println("Comportamento setado:", ok, errMsg)
		// },
		OnReading: func(client dtos.R2000ClientIface, reading dtos.ReadingStruct) {
			fmt.Println("Tag lida:", reading)
		},
	}

	client, err := clients.NewR2000Client("teste", "COM7", callbacks)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// loop infinito enviando comandos

	// aqui você chama uma função que escreve na serial
	// Exemplo (precisa existir no seu pacote connection/utils):
	// connection.SendCommand(client.Port, enums.GET_READER_TEMPERATURE)
	dto := &dtos.RealtimeDto{
		Antennas:     []int{0}, // antenas válidas 0x00–0x03
		Repeat:       3,
		DwellS:       1.1,
		SwitchDelayS: 0.005,
	}
	err = client.StartRealtime(dto)
	if err != nil {
		panic(err)
	}
	fmt.Println("Realtime iniciado")

	// deixa rodar 10 segundos
	time.Sleep(10 * time.Second)

	// para realtime
	client.StopRealtime()
	fmt.Println("Realtime parado")

}
