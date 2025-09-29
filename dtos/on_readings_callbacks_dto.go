package dtos

// Structure das possíveis callbacks do módulo.
// dtos/callbacks.go
type OnReadingCallbacks struct {
	OnReadingError       func(client R2000ClientIface, msg string)
	OnGetDrmStatus       func(client R2000ClientIface, status string)
	OnFirmware           func(client R2000ClientIface, version float64)
	OnTemperature        func(client R2000ClientIface, temperature int)
	OnGetWorkAntenna     func(client R2000ClientIface, antennas string)
	OnGetOutputPower     func(client R2000ClientIface, powers map[string]int)
	OnSetDrm             func(client R2000ClientIface, ok bool, errMsg string)
	OnSetWorkAntenna     func(client R2000ClientIface, ok bool, errMsg string)
	OnSetOutputPower     func(client R2000ClientIface, ok bool, errMsg string)
	OnSetBuzzerBehavior  func(client R2000ClientIface, ok bool, errMsg string)
	OnSetFrequencyRegion func(client R2000ClientIface, ok bool, errMsg string)
	OnReading            func(client R2000ClientIface, reading ReadingStruct)
	OnGetFrequencyRegion func(client R2000ClientIface, region string, frequency1, frequency2 float64, errMsg error)
}
