package clients

import (
	"sync"

	"github.com/viniggj2005/r2000-go/connection"
	"github.com/viniggj2005/r2000-go/utils"
	"go.bug.st/serial"
)

type OnReadingCallbacks struct {
	OnFirmware           func(client *R2000Client, version float64)
	OnTemperature        func(client *R2000Client, temperature int)
	OnReadingError       func(client *R2000Client, msg string)
	OnGetDrmStatus       func(client *R2000Client, status string)
	OnReading            func(client *R2000Client, reading utils.ReadingStruct)
	OnSetDrm             func(client *R2000Client, ok bool, errMsg *[]byte)
	OnGetWorkAntenna     func(client *R2000Client, antennas string)
	OnGetOutputPower     func(client *R2000Client, powers map[string]int)
	OnSetWorkAntenna     func(client *R2000Client, ok bool, errMsg *[]byte)
	OnSetOutputPower     func(client *R2000Client, ok bool, errMsg *[]byte)
	OnGetFrequencyRegion func(client *R2000Client, region string, f1, f2 float64, errMsg error)
	OnSetBuzzerBehavior  func(client *R2000Client, ok bool, errMsg *[]byte)
	OnSetFrequencyRegion func(client *R2000Client, ok bool, errMsg *[]byte)
}

// Estrutura principal do cliente
type R2000Client struct {
	Name       string      // Nome do cliente (identificação lógica)
	Port       serial.Port // Porta serial aberta
	Online     bool        // Status de conexão
	FrameQueue chan []byte // Canal para enfileirar frames recebidos
	//TODO criar uma struc para o OnFrame com a tratativa do tipo de cada callback
	Callbacks OnReadingCallbacks // Callback para processar cada frame recebido
	stopChan  chan struct{}      // Canal usado para sinalizar parada
	wg        sync.WaitGroup     // WaitGroup para aguardar goroutines terminarem
}

// Construtor do cliente: abre a porta, inicia listener e processador
func NewR2000Client(name, portName string, onFrame OnReadingCallbacks) (*R2000Client, error) {
	// Abre a porta serial com a função já existente no seu pacote connection
	portHandle, err := connection.OpenSerialConnection(portName)
	if err != nil {
		return nil, err
	}

	// Instancia o cliente
	client := &R2000Client{
		Name:       name,
		Port:       portHandle,
		Online:     true,
		FrameQueue: make(chan []byte, 100), // canal bufferizado para frames
		Callbacks:  onFrame,                // callback fornecido externamente
		stopChan:   make(chan struct{}),    // inicializa canal de parada
	}

	// Usa a função ListenSerial já criada no seu pacote connection
	// Cada vez que chegam bytes, eles são jogados na FrameQueue
	connection.ListenSerial(portHandle, func(data []byte) {
		select {
		case client.FrameQueue <- data: // envia para fila
		default: // se fila cheia, descarta
		}
	})

	// Inicia a goroutine que processa os frames
	client.wg.Add(1)
	go client.processFrames()

	return client, nil
}

// Fecha o cliente: sinaliza stopChan, aguarda goroutines e fecha a porta
func (c *R2000Client) Close() error {
	close(c.stopChan)     // avisa para as goroutines pararem
	c.wg.Wait()           // espera todas terminarem
	return c.Port.Close() // fecha a porta serial
}

// Goroutine que consome frames da fila e chama o callback OnFrame
func (c *R2000Client) processFrames() {
	defer c.wg.Done()
	for {
		select {
		case <-c.stopChan:
			return
		case frame := <-c.FrameQueue:
			// >>> aqui entra o dispatcher que você já criou
			ProcessR2000Frames(c, frame)
		}
	}
}
