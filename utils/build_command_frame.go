package utils

import (
	"github.com/viniggj2005/r2000-go/dtos"
	"github.com/viniggj2005/r2000-go/enums"
)

func ExtractParams(frame []byte) []byte {
	return frame[4 : len(frame)-1]
}

func ExtractReading(frame []byte) []byte {
	return frame[4 : len(frame)-1]
}

func EncodeParamsToBytes(params int) []byte {
	return []byte{byte(params)}
}

func BuildCommandFrame(obj dtos.BuildFrame) []byte {
	middle := []byte{}
	middle = append(middle, EncodeParamsToBytes(len(obj.Params)+3)...)
	middle = append(middle, 0x01)
	middle = append(middle, byte(obj.Command))
	middle = append(middle, obj.Params...)

	fullContent := append([]byte{byte(enums.HEADER)}, middle...)
	checksum := byte(CalculateChecksum(fullContent))
	frame := append(
		append([]byte{byte(enums.HEADER)}, middle...),
		checksum,
	)
	return frame
}
