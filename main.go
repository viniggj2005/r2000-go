package main

import (
	"fmt"

	"github.com/viniggj2005/r2000-go/dtos"
	"github.com/viniggj2005/r2000-go/enums"
	"github.com/viniggj2005/r2000-go/utils"
)

// "github.com/viniggj2005/r2000-go/mqtt"

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Erro ao carregar env .env file: %s", err)
	// }
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
	teste := utils.BuildCommandFrame(dtos.BuildFrame{Command: enums.SET_FREQUENCY_REGION, Params: []byte{byte(enums.FCC), byte(enums.MHZ_902_00), byte(enums.MHZ_928_00)}})
	fmt.Printf("%x\n", teste)
	//##################################################################

	// callbacks := dtos.OnReadingCallbacks{
	// 	OnFirmware: func(client dtos.R2000ClientIface, version float64) {
	// 		fmt.Println("Firmware version:", version)
	// 	},
	// 	OnTemperature: func(client dtos.R2000ClientIface, temperature int) {
	// 		fmt.Println("Temperature:", temperature, "cliente:", client.GetName())
	// 	},
	// 	OnGetDrmStatus: func(client dtos.R2000ClientIface, status string) {
	// 		fmt.Println("DRM Status:", status)
	// 	},
	// 	OnGetFrequencyRegion: func(client dtos.R2000ClientIface, region string, frequency1, frequency2 float64, err error) {
	// 		fmt.Println("region:", region, "f1:", frequency1, "f2:", frequency2)
	// 	},
	// 	OnGetOutputPower: func(client dtos.R2000ClientIface, power map[string]int) {
	// 		for chave, valor := range power {
	// 			fmt.Println("antenas:", chave, "potência :", valor, "Dbm")
	// 		}
	// 	},
	// 	OnGetWorkAntenna: func(client dtos.R2000ClientIface, antenna string) {
	// 		fmt.Println("region:", antenna)
	// 	},
	// 	OnSetBuzzerBehavior: func(client dtos.R2000ClientIface, ok bool, errMsg string) {
	// 		fmt.Println("Comportamento setado:", ok)
	// 	},
	// 	OnSetFrequencyRegion: func(client dtos.R2000ClientIface, ok bool, errMsg string) {
	// 		fmt.Println("Comportamento setado:", ok)
	// 	},
	// 	OnSetDrm: func(client dtos.R2000ClientIface, ok bool, errMsg string) {
	// 		fmt.Println("Comportamento setado:", ok)
	// 	},
	// 	OnSetOutputPower: func(client dtos.R2000ClientIface, ok bool, errMsg string) {
	// 		fmt.Println("Comportamento setado:", ok)
	// 	},
	// 	// OnSetWorkAntenna: func(client dtos.R2000ClientIface, ok bool, errMsg string) {
	// 	// 	fmt.Println("Comportamento setado:", ok, errMsg)
	// 	// },
	// 	OnReading: func(client dtos.R2000ClientIface, reading dtos.ReadingStruct) {
	// 		fmt.Println("Tag lida:", reading)
	// 	},
	// }

	// client, err := clients.NewR2000Client("teste", "COM7", callbacks)
	// if err != nil {
	// 	panic(err)
	// }
	// defer client.Close()

	// // loop infinito enviando comandos

	// // aqui você chama uma função que escreve na serial
	// // Exemplo (precisa existir no seu pacote connection/utils):
	// // connection.SendCommand(client.Port, enums.GET_READER_TEMPERATURE)
	// watcher := &clients.TimeWatcher{
	// 	TimeWatcherRun:      true,
	// 	TimeWatcherStopChan: make(chan struct{}),
	// }
	// watcher.StartTemperatureWatcher([]*clients.R2000Client{client}, 2)

	// time.Sleep(6 * time.Second)
	// watcher.StopTemperatureWatcher()

	// time.Sleep(10 * time.Second)

}
