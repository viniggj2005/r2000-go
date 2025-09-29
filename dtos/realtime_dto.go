package dtos

// RealtimeDto carrega a configuração para o modo realtime do R2000.
type RealtimeDto struct {
	Antennas     []int   // lista de antenas (0x00–0x03)
	Repeat       int     // número de repetições
	DwellS       float64 // tempo que cada antena fica ativa (segundos)
	SwitchDelayS float64 // delay entre troca de antenas (segundos)
}
