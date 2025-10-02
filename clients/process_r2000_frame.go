package clients

import (
	"fmt"

	"github.com/viniggj2005/r2000-go/enums"
	"github.com/viniggj2005/r2000-go/utils"
)

// Função responásvel por tratar os frames retornados do módulo pelo canal serial.
func ProcessR2000Frames(client *R2000Client, frame []byte) {
	validFrame := utils.ValidateFrame(frame)
	if !validFrame {
		return
	}

	frameCommand := frame[3]
	switch frameCommand {
	case byte(enums.GET_READER_TEMPERATURE):
		callback := client.Callbacks.OnTemperature
		if callback != nil {
			response := utils.OnGetTemperature(frame)
			callback(client, response)
		}
	case byte(enums.GET_FIRMWARE_VERSION):
		callback := client.Callbacks.OnFirmware
		if callback != nil {
			response := utils.OnGetFirmwareVersion(frame)
			callback(client, response)
		}
	case byte(enums.SET_BEEPER_MODE):
		callback := client.Callbacks.OnSetBuzzerBehavior
		if callback != nil {
			response, err := utils.OnSetMessage(frame)
			callback(client, response, err)
		}
	case byte(enums.GET_RF_POWER):
		callback := client.Callbacks.OnGetOutputPower
		if callback != nil {
			response := utils.OnGetOutPutPower(frame)
			callback(client, response)
		}
	case byte(enums.SET_TEMPORARY_OUTPUT_POWER):
		callback := client.Callbacks.OnSetOutputPower
		if callback != nil {
			response, err := utils.OnSetMessage(frame)
			callback(client, response, err)
		}
	case byte(enums.SET_RF_POWER):
		callback := client.Callbacks.OnSetOutputPower
		if callback != nil {
			response, err := utils.OnSetMessage(frame)
			callback(client, response, err)
		}
	case byte(enums.GET_WORK_ANTENNA):
		callback := client.Callbacks.OnGetWorkAntenna
		if callback != nil {
			response := utils.OnGetWorkAntenna(frame)
			callback(client, response)
		}
	case byte(enums.GET_FREQUENCY_REGION):
		callback := client.Callbacks.OnGetFrequencyRegion
		if callback != nil {
			response, f1, f2, err := utils.OnGetFrequencyRegion(frame)
			callback(client, response, f1, f2, err)
		}
	case byte(enums.SET_WORK_ANTENNA):
		callback := client.Callbacks.OnSetWorkAntenna
		if callback != nil {
			response, err := utils.OnSetMessage(frame)
			callback(client, response, err)
		}
	case byte(enums.SET_FREQUENCY_REGION):
		callback := client.Callbacks.OnSetFrequencyRegion
		if callback != nil {
			response, err := utils.OnSetMessage(frame)
			callback(client, response, err)
		}
	case byte(enums.GET_DRM_STATUS):
		callback := client.Callbacks.OnGetDrmStatus
		if callback != nil {
			response := utils.OnGetDrmStatus(frame)
			callback(client, response)
		}
	case byte(enums.SET_DRM):
		callback := client.Callbacks.OnSetDrm
		if callback != nil {
			response, err := utils.OnSetMessage(frame)
			callback(client, response, err)
		}

	case byte(enums.FAST_SWITCH_ANT_INVENTORY), byte(enums.REAL_TIME_INVENTORY):
		length := frame[1]

		if length == 0x05 && frame[5] == 0x22 {
			if cb := client.Callbacks.OnReadingError; cb != nil {
				antenna := frame[4] + 1
				cb(client, fmt.Sprintf("Antena %d ausente ou mal conectada.", antenna))
			}
		}

		if length >= 0x0F {
			if cb := client.Callbacks.OnReading; cb != nil {
				resp := utils.OnReading(frame)
				cb(client, resp)
			}
		}
	case byte(enums.RESET):
		if len(frame) >= 2 {
			fmt.Printf("Reset executado, código de resposta: 0x%X\n", frame[len(frame)-2])
		} else {
			fmt.Println("Frame inválido recebido no RESET")
		}
	default:
	}
}
