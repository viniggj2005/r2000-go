package enums

type R2000ErrorsEnum byte

const (
	FAIL                                            R2000ErrorsEnum = 0x11
	SUCCESS                                         R2000ErrorsEnum = 0x10
	CW_ON_ERROR                                     R2000ErrorsEnum = 0x21
	NO_TAG_ERROR                                    R2000ErrorsEnum = 0x36
	PLL_LOCK_FAIL                                   R2000ErrorsEnum = 0x52
	TAG_LOCK_ERROR                                  R2000ErrorsEnum = 0x34
	TAG_KILL_ERROR                                  R2000ErrorsEnum = 0x35
	TAG_READ_ERROR                                  R2000ErrorsEnum = 0x32
	MCU_RESET_ERROR                                 R2000ErrorsEnum = 0x20
	TAG_WRITE_ERROR                                 R2000ErrorsEnum = 0x33
	READ_FLASH_ERROR                                R2000ErrorsEnum = 0x24
	WRITE_FLASH_ERROR                               R2000ErrorsEnum = 0x23
	PARAMETER_INVALID                               R2000ErrorsEnum = 0x41
	TAG_INVENTORY_ERROR                             R2000ErrorsEnum = 0x31
	OUTPUT_POWER_TOO_LOW                            R2000ErrorsEnum = 0x57
	BUFFER_IS_EMPTY_ERROR                           R2000ErrorsEnum = 0x38
	ANTENNA_MISSING_ERROR                           R2000ErrorsEnum = 0x22
	SET_OUTPUT_POWER_ERROR                          R2000ErrorsEnum = 0x25
	NXP_CUSTOM_COMMAND_FAIL                         R2000ErrorsEnum = 0x3C
	ACCESS_OR_PASSWORD_ERROR                        R2000ErrorsEnum = 0x40
	RF_CHIP_FAIL_TO_RESPONSE                        R2000ErrorsEnum = 0x53
	SPECTRUM_REGULATION_ERROR                       R2000ErrorsEnum = 0x56
	FAIL_TO_GET_RN16_FROM_TAG                       R2000ErrorsEnum = 0x50
	PARAMETER_INVALID_DRM_MODE                      R2000ErrorsEnum = 0x51
	INVENTORY_OK_BUT_ACCESS_FAIL                    R2000ErrorsEnum = 0x37
	COPYRIGHT_AUTHENTICATION_FAIL                   R2000ErrorsEnum = 0x55
	PARAMETER_EPC_MATCH_LEN_ERROR                   R2000ErrorsEnum = 0x4D
	FAIL_TO_GET_RF_PORT_RETURN_LOSS                 R2000ErrorsEnum = 0xEE
	PARAMETER_INVALID_EPC_MATCH_MODE                R2000ErrorsEnum = 0x4E
	PARAMETER_READER_ADDRESS_INVALID                R2000ErrorsEnum = 0x46
	PARAMETER_EPC_MATCH_LEN_TOO_LONG                R2000ErrorsEnum = 0x4C
	PARAMETER_INVALID_FREQUENCY_RANGE               R2000ErrorsEnum = 0x4F
	PARAMETER_INVALID_WORDCNT_TOO_LONG              R2000ErrorsEnum = 0x42
	PARAMETER_BEEPER_MODE_OUT_OF_RANGE              R2000ErrorsEnum = 0x4B
	FAIL_TO_ACHIEVE_DESIRED_OUTPUT_POWER            R2000ErrorsEnum = 0x54
	PARAMETER_INVALID_MEMBANK_OUT_OF_RANGE          R2000ErrorsEnum = 0x43
	PARAMETER_INVALID_BAUDRATE_OUT_OF_RANGE         R2000ErrorsEnum = 0x4A
	PARAMETER_INVALID_ANTENNA_ID_OUT_OF_RANGE       R2000ErrorsEnum = 0x47
	PARAMETER_INVALID_LOCK_REGION_OUT_OF_RANGE      R2000ErrorsEnum = 0x44
	PARAMETER_INVALID_LOCK_ACTION_OUT_OF_RANGE      R2000ErrorsEnum = 0x45
	PARAMETER_INVALID_OUTPUT_POWER_OUT_OF_RANGE     R2000ErrorsEnum = 0x48
	PARAMETER_INVALID_FREQUENCY_REGION_OUT_OF_RANGE R2000ErrorsEnum = 0x49
)

var Errors = struct {
	FAIL                                            []byte
	SUCCESS                                         []byte
	CW_ON_ERROR                                     []byte
	NO_TAG_ERROR                                    []byte
	PLL_LOCK_FAIL                                   []byte
	TAG_LOCK_ERROR                                  []byte
	TAG_KILL_ERROR                                  []byte
	TAG_READ_ERROR                                  []byte
	MCU_RESET_ERROR                                 []byte
	TAG_WRITE_ERROR                                 []byte
	READ_FLASH_ERROR                                []byte
	WRITE_FLASH_ERROR                               []byte
	PARAMETER_INVALID                               []byte
	TAG_INVENTORY_ERROR                             []byte
	OUTPUT_POWER_TOO_LOW                            []byte
	BUFFER_IS_EMPTY_ERROR                           []byte
	ANTENNA_MISSING_ERROR                           []byte
	SET_OUTPUT_POWER_ERROR                          []byte
	NXP_CUSTOM_COMMAND_FAIL                         []byte
	ACCESS_OR_PASSWORD_ERROR                        []byte
	RF_CHIP_FAIL_TO_RESPONSE                        []byte
	SPECTRUM_REGULATION_ERROR                       []byte
	FAIL_TO_GET_RN16_FROM_TAG                       []byte
	PARAMETER_INVALID_DRM_MODE                      []byte
	INVENTORY_OK_BUT_ACCESS_FAIL                    []byte
	COPYRIGHT_AUTHENTICATION_FAIL                   []byte
	PARAMETER_EPC_MATCH_LEN_ERROR                   []byte
	FAIL_TO_GET_RF_PORT_RETURN_LOSS                 []byte
	PARAMETER_INVALID_EPC_MATCH_MODE                []byte
	PARAMETER_READER_ADDRESS_INVALID                []byte
	PARAMETER_EPC_MATCH_LEN_TOO_LONG                []byte
	PARAMETER_INVALID_FREQUENCY_RANGE               []byte
	PARAMETER_INVALID_WORDCNT_TOO_LONG              []byte
	PARAMETER_BEEPER_MODE_OUT_OF_RANGE              []byte
	FAIL_TO_ACHIEVE_DESIRED_OUTPUT_POWER            []byte
	PARAMETER_INVALID_MEMBANK_OUT_OF_RANGE          []byte
	PARAMETER_INVALID_BAUDRATE_OUT_OF_RANGE         []byte
	PARAMETER_INVALID_ANTENNA_ID_OUT_OF_RANGE       []byte
	PARAMETER_INVALID_LOCK_REGION_OUT_OF_RANGE      []byte
	PARAMETER_INVALID_LOCK_ACTION_OUT_OF_RANGE      []byte
	PARAMETER_INVALID_OUTPUT_POWER_OUT_OF_RANGE     []byte
	PARAMETER_INVALID_FREQUENCY_REGION_OUT_OF_RANGE []byte
}{
	FAIL:                                            []byte{byte(FAIL)},
	SUCCESS:                                         []byte{byte(SUCCESS)},
	CW_ON_ERROR:                                     []byte{byte(CW_ON_ERROR)},
	NO_TAG_ERROR:                                    []byte{byte(NO_TAG_ERROR)},
	PLL_LOCK_FAIL:                                   []byte{byte(PLL_LOCK_FAIL)},
	TAG_LOCK_ERROR:                                  []byte{byte(TAG_LOCK_ERROR)},
	TAG_KILL_ERROR:                                  []byte{byte(TAG_KILL_ERROR)},
	TAG_READ_ERROR:                                  []byte{byte(TAG_READ_ERROR)},
	MCU_RESET_ERROR:                                 []byte{byte(MCU_RESET_ERROR)},
	TAG_WRITE_ERROR:                                 []byte{byte(TAG_WRITE_ERROR)},
	READ_FLASH_ERROR:                                []byte{byte(READ_FLASH_ERROR)},
	WRITE_FLASH_ERROR:                               []byte{byte(WRITE_FLASH_ERROR)},
	PARAMETER_INVALID:                               []byte{byte(PARAMETER_INVALID)},
	TAG_INVENTORY_ERROR:                             []byte{byte(TAG_INVENTORY_ERROR)},
	OUTPUT_POWER_TOO_LOW:                            []byte{byte(OUTPUT_POWER_TOO_LOW)},
	BUFFER_IS_EMPTY_ERROR:                           []byte{byte(BUFFER_IS_EMPTY_ERROR)},
	ANTENNA_MISSING_ERROR:                           []byte{byte(ANTENNA_MISSING_ERROR)},
	SET_OUTPUT_POWER_ERROR:                          []byte{byte(SET_OUTPUT_POWER_ERROR)},
	NXP_CUSTOM_COMMAND_FAIL:                         []byte{byte(NXP_CUSTOM_COMMAND_FAIL)},
	ACCESS_OR_PASSWORD_ERROR:                        []byte{byte(ACCESS_OR_PASSWORD_ERROR)},
	RF_CHIP_FAIL_TO_RESPONSE:                        []byte{byte(RF_CHIP_FAIL_TO_RESPONSE)},
	SPECTRUM_REGULATION_ERROR:                       []byte{byte(SPECTRUM_REGULATION_ERROR)},
	FAIL_TO_GET_RN16_FROM_TAG:                       []byte{byte(FAIL_TO_GET_RN16_FROM_TAG)},
	PARAMETER_INVALID_DRM_MODE:                      []byte{byte(PARAMETER_INVALID_DRM_MODE)},
	INVENTORY_OK_BUT_ACCESS_FAIL:                    []byte{byte(INVENTORY_OK_BUT_ACCESS_FAIL)},
	COPYRIGHT_AUTHENTICATION_FAIL:                   []byte{byte(COPYRIGHT_AUTHENTICATION_FAIL)},
	PARAMETER_EPC_MATCH_LEN_ERROR:                   []byte{byte(PARAMETER_EPC_MATCH_LEN_ERROR)},
	FAIL_TO_GET_RF_PORT_RETURN_LOSS:                 []byte{byte(FAIL_TO_GET_RF_PORT_RETURN_LOSS)},
	PARAMETER_INVALID_EPC_MATCH_MODE:                []byte{byte(PARAMETER_INVALID_EPC_MATCH_MODE)},
	PARAMETER_READER_ADDRESS_INVALID:                []byte{byte(PARAMETER_READER_ADDRESS_INVALID)},
	PARAMETER_EPC_MATCH_LEN_TOO_LONG:                []byte{byte(PARAMETER_EPC_MATCH_LEN_TOO_LONG)},
	PARAMETER_INVALID_FREQUENCY_RANGE:               []byte{byte(PARAMETER_INVALID_FREQUENCY_RANGE)},
	PARAMETER_INVALID_WORDCNT_TOO_LONG:              []byte{byte(PARAMETER_INVALID_WORDCNT_TOO_LONG)},
	PARAMETER_BEEPER_MODE_OUT_OF_RANGE:              []byte{byte(PARAMETER_BEEPER_MODE_OUT_OF_RANGE)},
	FAIL_TO_ACHIEVE_DESIRED_OUTPUT_POWER:            []byte{byte(FAIL_TO_ACHIEVE_DESIRED_OUTPUT_POWER)},
	PARAMETER_INVALID_MEMBANK_OUT_OF_RANGE:          []byte{byte(PARAMETER_INVALID_MEMBANK_OUT_OF_RANGE)},
	PARAMETER_INVALID_BAUDRATE_OUT_OF_RANGE:         []byte{byte(PARAMETER_INVALID_BAUDRATE_OUT_OF_RANGE)},
	PARAMETER_INVALID_ANTENNA_ID_OUT_OF_RANGE:       []byte{byte(PARAMETER_INVALID_ANTENNA_ID_OUT_OF_RANGE)},
	PARAMETER_INVALID_LOCK_REGION_OUT_OF_RANGE:      []byte{byte(PARAMETER_INVALID_LOCK_REGION_OUT_OF_RANGE)},
	PARAMETER_INVALID_LOCK_ACTION_OUT_OF_RANGE:      []byte{byte(PARAMETER_INVALID_LOCK_ACTION_OUT_OF_RANGE)},
	PARAMETER_INVALID_OUTPUT_POWER_OUT_OF_RANGE:     []byte{byte(PARAMETER_INVALID_OUTPUT_POWER_OUT_OF_RANGE)},
	PARAMETER_INVALID_FREQUENCY_REGION_OUT_OF_RANGE: []byte{byte(PARAMETER_INVALID_FREQUENCY_REGION_OUT_OF_RANGE)},
}
