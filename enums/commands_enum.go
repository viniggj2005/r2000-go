package enums

type R2000CommandsEnum byte

const (
	SET_TEMPORARY_OUTPUT_POWER R2000CommandsEnum = 0x66
	GET_DRM_STATUS             R2000CommandsEnum = 0x7D
	SET_DRM                    R2000CommandsEnum = 0x7C
	RESET                      R2000CommandsEnum = 0x70
	GET_FIRMWARE_VERSION       R2000CommandsEnum = 0x72
	SET_WORK_ANTENNA           R2000CommandsEnum = 0x74
	GET_WORK_ANTENNA           R2000CommandsEnum = 0x75
	SET_RF_POWER               R2000CommandsEnum = 0x76
	GET_RF_POWER               R2000CommandsEnum = 0x77
	SET_FREQUENCY_REGION       R2000CommandsEnum = 0x78
	GET_FREQUENCY_REGION       R2000CommandsEnum = 0x79
	SET_BEEPER_MODE            R2000CommandsEnum = 0x7A
	GET_READER_TEMPERATURE     R2000CommandsEnum = 0x7B
	REAL_TIME_INVENTORY        R2000CommandsEnum = 0x89
	FAST_SWITCH_ANT_INVENTORY  R2000CommandsEnum = 0x8A
)

var Commands = struct {
	RESET                      []byte
	SET_DRM                    []byte
	SET_RF_POWER               []byte
	GET_RF_POWER               []byte
	GET_DRM_STATUS             []byte
	SET_BEEPER_MODE            []byte
	SET_WORK_ANTENNA           []byte
	GET_WORK_ANTENNA           []byte
	REAL_TIME_INVENTORY        []byte
	GET_FIRMWARE_VERSION       []byte
	SET_FREQUENCY_REGION       []byte
	GET_FREQUENCY_REGION       []byte
	GET_READER_TEMPERATURE     []byte
	FAST_SWITCH_ANT_INVENTORY  []byte
	SET_TEMPORARY_OUTPUT_POWER []byte
}{
	RESET:                      []byte{byte(RESET)},
	SET_DRM:                    []byte{byte(SET_DRM)},
	SET_RF_POWER:               []byte{byte(SET_RF_POWER)},
	GET_RF_POWER:               []byte{byte(GET_RF_POWER)},
	GET_DRM_STATUS:             []byte{byte(GET_DRM_STATUS)},
	SET_BEEPER_MODE:            []byte{byte(SET_BEEPER_MODE)},
	SET_WORK_ANTENNA:           []byte{byte(SET_WORK_ANTENNA)},
	GET_WORK_ANTENNA:           []byte{byte(GET_WORK_ANTENNA)},
	REAL_TIME_INVENTORY:        []byte{byte(REAL_TIME_INVENTORY)},
	SET_FREQUENCY_REGION:       []byte{byte(SET_FREQUENCY_REGION)},
	GET_FIRMWARE_VERSION:       []byte{byte(GET_FIRMWARE_VERSION)},
	GET_FREQUENCY_REGION:       []byte{byte(GET_FREQUENCY_REGION)},
	GET_READER_TEMPERATURE:     []byte{byte(GET_READER_TEMPERATURE)},
	FAST_SWITCH_ANT_INVENTORY:  []byte{byte(FAST_SWITCH_ANT_INVENTORY)},
	SET_TEMPORARY_OUTPUT_POWER: []byte{byte(SET_TEMPORARY_OUTPUT_POWER)},
}
