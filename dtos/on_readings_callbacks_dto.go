package dtos

// R2000ClientIface deve existir no seu projeto.
// Aqui só referenciamos para não quebrar assinaturas.
type OnReadingCallbacks struct {
	// Novos campos opcionais para observabilidade
	OnError    func(error)  // erros de I/O e parsing
	OnRawFrame func([]byte) // frame bruto antes do dispatcher

	OnSetDrm             func(client R2000ClientIface, ok bool, errMsg string)
	OnReading            func(client R2000ClientIface, reading ReadingStruct)
	OnFirmware           func(client R2000ClientIface, version float64)
	OnTemperature        func(client R2000ClientIface, temperature int)
	OnGetDrmStatus       func(client R2000ClientIface, status string)
	OnReadingError       func(client R2000ClientIface, msg string)
	OnGetWorkAntenna     func(client R2000ClientIface, antennas string)
	OnGetOutputPower     func(client R2000ClientIface, powers map[string]int)
	OnSetWorkAntenna     func(client R2000ClientIface, ok bool, errMsg string)
	OnSetOutputPower     func(client R2000ClientIface, ok bool, errMsg string)
	OnSetBuzzerBehavior  func(client R2000ClientIface, ok bool, errMsg string)
	OnSetFrequencyRegion func(client R2000ClientIface, ok bool, errMsg string)
	OnGetFrequencyRegion func(client R2000ClientIface, region string, start, end float64, errMsg error)
}
