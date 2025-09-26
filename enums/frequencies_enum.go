package enums

import (
	"fmt"
	"strconv"
	"strings"
)

type R2000FrequencyEnum byte

const (
	MHZ_865_00 R2000FrequencyEnum = 0x00
	MHZ_865_50 R2000FrequencyEnum = 0x01
	MHZ_866_00 R2000FrequencyEnum = 0x02
	MHZ_866_50 R2000FrequencyEnum = 0x03
	MHZ_867_00 R2000FrequencyEnum = 0x04
	MHZ_867_50 R2000FrequencyEnum = 0x05
	MHZ_868_00 R2000FrequencyEnum = 0x06
	MHZ_902_00 R2000FrequencyEnum = 0x07
	MHZ_902_50 R2000FrequencyEnum = 0x08
	MHZ_903_00 R2000FrequencyEnum = 0x09
	MHZ_903_50 R2000FrequencyEnum = 0x0A
	MHZ_904_00 R2000FrequencyEnum = 0x0B
	MHZ_904_50 R2000FrequencyEnum = 0x0C
	MHZ_905_00 R2000FrequencyEnum = 0x0D
	MHZ_905_50 R2000FrequencyEnum = 0x0E
	MHZ_906_00 R2000FrequencyEnum = 0x0F
	MHZ_906_50 R2000FrequencyEnum = 0x10
	MHZ_907_00 R2000FrequencyEnum = 0x11
	MHZ_907_50 R2000FrequencyEnum = 0x12
	MHZ_908_00 R2000FrequencyEnum = 0x13
	MHZ_908_50 R2000FrequencyEnum = 0x14
	MHZ_909_00 R2000FrequencyEnum = 0x15
	MHZ_909_50 R2000FrequencyEnum = 0x16
	MHZ_910_00 R2000FrequencyEnum = 0x17
	MHZ_910_50 R2000FrequencyEnum = 0x18
	MHZ_911_00 R2000FrequencyEnum = 0x19
	MHZ_911_50 R2000FrequencyEnum = 0x1A
	MHZ_912_00 R2000FrequencyEnum = 0x1B
	MHZ_912_50 R2000FrequencyEnum = 0x1C
	MHZ_913_00 R2000FrequencyEnum = 0x1D
	MHZ_913_50 R2000FrequencyEnum = 0x1E
	MHZ_914_00 R2000FrequencyEnum = 0x1F
	MHZ_914_50 R2000FrequencyEnum = 0x20
	MHZ_915_00 R2000FrequencyEnum = 0x21
	MHZ_915_50 R2000FrequencyEnum = 0x22
	MHZ_916_00 R2000FrequencyEnum = 0x23
	MHZ_916_50 R2000FrequencyEnum = 0x24
	MHZ_917_00 R2000FrequencyEnum = 0x25
	MHZ_917_50 R2000FrequencyEnum = 0x26
	MHZ_918_00 R2000FrequencyEnum = 0x27
	MHZ_918_50 R2000FrequencyEnum = 0x28
	MHZ_919_00 R2000FrequencyEnum = 0x29
	MHZ_919_50 R2000FrequencyEnum = 0x2A
	MHZ_920_00 R2000FrequencyEnum = 0x2B
	MHZ_920_50 R2000FrequencyEnum = 0x2C
	MHZ_921_00 R2000FrequencyEnum = 0x2D
	MHZ_921_50 R2000FrequencyEnum = 0x2E
	MHZ_922_00 R2000FrequencyEnum = 0x2F
	MHZ_922_50 R2000FrequencyEnum = 0x30
	MHZ_923_00 R2000FrequencyEnum = 0x31
	MHZ_923_50 R2000FrequencyEnum = 0x32
	MHZ_924_00 R2000FrequencyEnum = 0x33
	MHZ_924_50 R2000FrequencyEnum = 0x34
	MHZ_925_00 R2000FrequencyEnum = 0x35
	MHZ_925_50 R2000FrequencyEnum = 0x36
	MHZ_926_00 R2000FrequencyEnum = 0x37
	MHZ_926_50 R2000FrequencyEnum = 0x38
	MHZ_927_00 R2000FrequencyEnum = 0x39
	MHZ_927_50 R2000FrequencyEnum = 0x3A
	MHZ_928_00 R2000FrequencyEnum = 0x3B
)

var freqNames = map[R2000FrequencyEnum]string{
	MHZ_865_00: "MHZ_865_00",
	MHZ_865_50: "MHZ_865_50",
	MHZ_866_00: "MHZ_866_00",
	MHZ_866_50: "MHZ_866_50",
	MHZ_867_00: "MHZ_867_00",
	MHZ_867_50: "MHZ_867_50",
	MHZ_868_00: "MHZ_868_00",
	MHZ_902_00: "MHZ_902_00",
	MHZ_902_50: "MHZ_902_50",
	MHZ_903_00: "MHZ_903_00",
	MHZ_903_50: "MHZ_903_50",
	MHZ_904_00: "MHZ_904_00",
	MHZ_904_50: "MHZ_904_50",
	MHZ_905_00: "MHZ_905_00",
	MHZ_905_50: "MHZ_905_50",
	MHZ_906_00: "MHZ_906_00",
	MHZ_906_50: "MHZ_906_50",
	MHZ_907_00: "MHZ_907_00",
	MHZ_907_50: "MHZ_907_50",
	MHZ_908_00: "MHZ_908_00",
	MHZ_908_50: "MHZ_908_50",
	MHZ_909_00: "MHZ_909_00",
	MHZ_909_50: "MHZ_909_50",
	MHZ_910_00: "MHZ_910_00",
	MHZ_910_50: "MHZ_910_50",
	MHZ_911_00: "MHZ_911_00",
	MHZ_911_50: "MHZ_911_50",
	MHZ_912_00: "MHZ_912_00",
	MHZ_912_50: "MHZ_912_50",
	MHZ_913_00: "MHZ_913_00",
	MHZ_913_50: "MHZ_913_50",
	MHZ_914_00: "MHZ_914_00",
	MHZ_914_50: "MHZ_914_50",
	MHZ_915_00: "MHZ_915_00",
	MHZ_915_50: "MHZ_915_50",
	MHZ_916_00: "MHZ_916_00",
	MHZ_916_50: "MHZ_916_50",
	MHZ_917_00: "MHZ_917_00",
	MHZ_917_50: "MHZ_917_50",
	MHZ_918_00: "MHZ_918_00",
	MHZ_918_50: "MHZ_918_50",
	MHZ_919_00: "MHZ_919_00",
	MHZ_919_50: "MHZ_919_50",
	MHZ_920_00: "MHZ_920_00",
	MHZ_920_50: "MHZ_920_50",
	MHZ_921_00: "MHZ_921_00",
	MHZ_921_50: "MHZ_921_50",
	MHZ_922_00: "MHZ_922_00",
	MHZ_922_50: "MHZ_922_50",
	MHZ_923_00: "MHZ_923_00",
	MHZ_923_50: "MHZ_923_50",
	MHZ_924_00: "MHZ_924_00",
	MHZ_924_50: "MHZ_924_50",
	MHZ_925_00: "MHZ_925_00",
	MHZ_925_50: "MHZ_925_50",
	MHZ_926_00: "MHZ_926_00",
	MHZ_926_50: "MHZ_926_50",
	MHZ_927_00: "MHZ_927_00",
	MHZ_927_50: "MHZ_927_50",
	MHZ_928_00: "MHZ_928_00",
}

func GetFrequency(val R2000FrequencyEnum) (float64, error) {
	name, ok := freqNames[val]
	if !ok {
		return 0, fmt.Errorf("frequência desconhecida: 0x%X", byte(val))
	}

	parts := strings.Split(name, "_")
	if len(parts) < 3 {
		return 0, fmt.Errorf("formato inválido: %s", name)
	}
	freqStr := parts[1] + "." + parts[2]

	return strconv.ParseFloat(freqStr, 64)
}
