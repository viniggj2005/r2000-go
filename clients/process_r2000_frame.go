package clients

import (
	"fmt"

	"github.com/viniggj2005/r2000-go/enums"
	"github.com/viniggj2005/r2000-go/utils"
)

// Tratamento de frames do protocolo R2000.
func ProcessR2000Frames(client *R2000Client, frame []byte) {
	if !utils.ValidateFrame(frame) {
		if cb := client.Callbacks.OnReadingError; cb != nil {
			cb(client, "frame inválido")
		}
		return
	}

	if len(frame) < 4 {
		if cb := client.Callbacks.OnReadingError; cb != nil {
			cb(client, "frame muito curto")
		}
		return
	}

	cmd := frame[3]

	switch cmd {
	case byte(enums.GET_READER_TEMPERATURE):
		if cb := client.Callbacks.OnTemperature; cb != nil {
			cb(client, utils.OnGetTemperature(frame))
		}

	case byte(enums.GET_FIRMWARE_VERSION):
		if cb := client.Callbacks.OnFirmware; cb != nil {
			cb(client, utils.OnGetFirmwareVersion(frame))
		}

	case byte(enums.SET_BEEPER_MODE):
		if cb := client.Callbacks.OnSetBuzzerBehavior; cb != nil {
			ok, err := utils.OnSetMessage(frame)
			cb(client, ok, err)
		}

	case byte(enums.GET_RF_POWER):
		if cb := client.Callbacks.OnGetOutputPower; cb != nil {
			cb(client, utils.OnGetOutPutPower(frame))
		}

	case byte(enums.SET_TEMPORARY_OUTPUT_POWER), byte(enums.SET_RF_POWER):
		if cb := client.Callbacks.OnSetOutputPower; cb != nil {
			ok, err := utils.OnSetMessage(frame)
			cb(client, ok, err)
		}

	case byte(enums.GET_WORK_ANTENNA):
		if cb := client.Callbacks.OnGetWorkAntenna; cb != nil {
			cb(client, utils.OnGetWorkAntenna(frame))
		}

	case byte(enums.GET_FREQUENCY_REGION):
		if cb := client.Callbacks.OnGetFrequencyRegion; cb != nil {
			region, f1, f2, err := utils.OnGetFrequencyRegion(frame)
			cb(client, region, f1, f2, err)
		}

	case byte(enums.SET_WORK_ANTENNA):
		if cb := client.Callbacks.OnSetWorkAntenna; cb != nil {
			ok, err := utils.OnSetMessage(frame)
			cb(client, ok, err)
		}

	case byte(enums.SET_FREQUENCY_REGION):
		if cb := client.Callbacks.OnSetFrequencyRegion; cb != nil {
			ok, err := utils.OnSetMessage(frame)
			cb(client, ok, err)
		}

	case byte(enums.GET_DRM_STATUS):
		if cb := client.Callbacks.OnGetDrmStatus; cb != nil {
			cb(client, utils.OnGetDrmStatus(frame))
		}

	case byte(enums.SET_DRM):
		if cb := client.Callbacks.OnSetDrm; cb != nil {
			ok, err := utils.OnSetMessage(frame)
			cb(client, ok, err)
		}

	case byte(enums.FAST_SWITCH_ANT_INVENTORY), byte(enums.REAL_TIME_INVENTORY):
		length := frame[1]

		if length == 0x05 && len(frame) >= 6 && frame[5] == 0x22 {
			if cb := client.Callbacks.OnReadingError; cb != nil {
				antenna := frame[4] + 1
				cb(client, fmt.Sprintf("Antena %d ausente ou mal conectada.", antenna))
			}
		}

		if length >= 0x0F {
			if cb := client.Callbacks.OnReading; cb != nil {
				cb(client, utils.OnReading(frame))
			}
		}

	case byte(enums.RESET):
		if len(frame) >= 2 {
			fmt.Printf("Reset executado, código de resposta: 0x%X\n", frame[len(frame)-2])
		} else {
			fmt.Println("Frame inválido no RESET")
		}

	default:
		// silencioso por padrão
	}
}
