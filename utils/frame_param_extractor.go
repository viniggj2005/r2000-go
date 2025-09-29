package utils

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/viniggj2005/r2000-go/dtos"
	"github.com/viniggj2005/r2000-go/enums"
)

func OnGetTemperature(frame []byte) int {
	params := ExtractParams(frame)
	var temperature int
	if params[0] == 0 {
		temperature = int(params[1]) * -1
	} else {
		temperature = int(params[1])
	}
	return temperature
}

func OnGetFirmwareVersion(frame []byte) float64 {
	params := ExtractParams(frame)
	str := fmt.Sprintf("%v.%v", params[0], params[1])
	version, _ := strconv.ParseFloat(str, 64)
	return version
}

func OnGetOutPutPower(frame []byte) map[string]int {
	params := ExtractParams(frame)
	power := make(map[string]int)
	if len(params) > 1 {
		for i, byteValue := range params {
			power[fmt.Sprintf("ant%d", i+1)] = int(byteValue)
		}
	} else {
		power["all"] = int(params[0])
	}
	return power
}

func OnGetWorkAntenna(frame []byte) string {
	params := ExtractParams(frame)
	var workAntenna string
	for _, byteValue := range params {
		workAntenna = fmt.Sprintf("Antenna%d", byteValue+1)
	}
	return workAntenna
}

func OnGetDrmStatus(frame []byte) string {
	params := ExtractParams(frame)
	var drmStatus string
	if params[0] == 1 {
		drmStatus = "On"
	} else {
		drmStatus = "Off"
	}
	return drmStatus
}
func OnGetFrequencyRegion(frame []byte) (string, float64, float64, error) {
	params := ExtractParams(frame)

	if params[0] != byte(enums.USER) {
		frequencyArray := make([]float64, 0, len(params)-1)

		for _, item := range params[1:] {
			freq, err := enums.GetFrequency(enums.R2000FrequencyEnum(item))
			if err != nil {
				return "", 0, 0, err
			}
			frequencyArray = append(frequencyArray, freq)
		}

		if len(frequencyArray) < 2 {
			return "", 0, 0, fmt.Errorf("frequências insuficientes")
		}

		return enums.R2000RegionsEnum(params[0]).String(),
			frequencyArray[0],
			frequencyArray[1],
			nil
	}
	return "USER", 0, 0, nil
}

func OnSetMessage(frame []byte) (bool, string) {
	params := ExtractParams(frame)
	var errorCode string
	var response bool
	if len(params) > 0 && params[0] == byte(enums.SUCCESS) {
		response = true
	} else {
		errorCode = string(params)
		response = false
	}
	return response, errorCode
}

func OnReading(frame []byte) dtos.ReadingStruct {
	reading := ExtractReading(frame)
	antenna := int(reading[0] & 0b11)
	pc := strings.ToUpper(hex.EncodeToString(reading[1:3]))
	epc := strings.ToUpper(hex.EncodeToString(reading[3:]))
	return dtos.ReadingStruct{
		Antenna: antenna,
		Pc:      pc,
		Epc:     epc,
	}
}
