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
