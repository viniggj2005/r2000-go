package enums

type R2000ErrorsEnum byte

const (
	FAIL                                            R2000ErrorsEnum = 0x11
	SUCCESS                                         R2000ErrorsEnum = 0x10
	CW_ON_ERROR                                     R2000ErrorsEnum = 0x21
	NO_TAG_ERROR                                    R2000ErrorsEnum = 0x36
	TAG_LOCK_ERROR                                  R2000ErrorsEnum = 0x34
	TAG_KILL_ERROR                                  R2000ErrorsEnum = 0x35
	TAG_READ_ERROR                                  R2000ErrorsEnum = 0x32
	MCU_RESET_ERROR                                 R2000ErrorsEnum = 0x20
	TAG_WRITE_ERROR                                 R2000ErrorsEnum = 0x33
	READ_FLASH_ERROR                                R2000ErrorsEnum = 0x24
	WRITE_FLASH_ERROR                               R2000ErrorsEnum = 0x23
	TAG_INVENTORY_ERROR                             R2000ErrorsEnum = 0x31
	ANTENNA_MISSING_ERROR                           R2000ErrorsEnum = 0x22
	SET_OUTPUT_POWER_ERROR                          R2000ErrorsEnum = 0x25
	INVENTORY_OK_BUT_ACCESS_FAIL                    R2000ErrorsEnum = 0x37
	BUFFER_IS_EMPTY_ERROR                           R2000ErrorsEnum = 0x38
	NXP_CUSTOM_COMMAND_FAIL                         R2000ErrorsEnum = 0x3C
	FAIL_TO_GET_RN16_FROM_TAG                       R2000ErrorsEnum = 0x50
	PARAMETER_INVALID_DRM_MODE                      R2000ErrorsEnum = 0x51
	PLL_LOCK_FAIL                                   R2000ErrorsEnum = 0x52
	RF_CHIP_FAIL_TO_RESPONSE                        R2000ErrorsEnum = 0x53
	FAIL_TO_ACHIEVE_DESIRED_OUTPUT_POWER            R2000ErrorsEnum = 0x54
	COPYRIGHT_AUTHENTICATION_FAIL                   R2000ErrorsEnum = 0x55
	SPECTRUM_REGULATION_ERROR                       R2000ErrorsEnum = 0x56
	OUTPUT_POWER_TOO_LOW                            R2000ErrorsEnum = 0x57
	FAIL_TO_GET_RF_PORT_RETURN_LOSS                 R2000ErrorsEnum = 0xEE
	ACCESS_OR_PASSWORD_ERROR                        R2000ErrorsEnum = 0x40
	PARAMETER_INVALID                               R2000ErrorsEnum = 0x41
	PARAMETER_INVALID_WORDCNT_TOO_LONG              R2000ErrorsEnum = 0x42
	PARAMETER_INVALID_MEMBANK_OUT_OF_RANGE          R2000ErrorsEnum = 0x43
	PARAMETER_INVALID_LOCK_REGION_OUT_OF_RANGE      R2000ErrorsEnum = 0x44
	PARAMETER_INVALID_LOCK_ACTION_OUT_OF_RANGE      R2000ErrorsEnum = 0x45
	PARAMETER_READER_ADDRESS_INVALID                R2000ErrorsEnum = 0x46
	PARAMETER_INVALID_ANTENNA_ID_OUT_OF_RANGE       R2000ErrorsEnum = 0x47
	PARAMETER_INVALID_OUTPUT_POWER_OUT_OF_RANGE     R2000ErrorsEnum = 0x48
	PARAMETER_INVALID_FREQUENCY_REGION_OUT_OF_RANGE R2000ErrorsEnum = 0x49
	PARAMETER_INVALID_BAUDRATE_OUT_OF_RANGE         R2000ErrorsEnum = 0x4A
	PARAMETER_BEEPER_MODE_OUT_OF_RANGE              R2000ErrorsEnum = 0x4B
	PARAMETER_EPC_MATCH_LEN_TOO_LONG                R2000ErrorsEnum = 0x4C
	PARAMETER_EPC_MATCH_LEN_ERROR                   R2000ErrorsEnum = 0x4D
	PARAMETER_INVALID_EPC_MATCH_MODE                R2000ErrorsEnum = 0x4E
	PARAMETER_INVALID_FREQUENCY_RANGE               R2000ErrorsEnum = 0x4F
)
