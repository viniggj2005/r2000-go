package enums

type R2000FastSwitchInventoryEnum byte

const (
	ANTENNA1 R2000FastSwitchInventoryEnum = 0x00
	ANTENNA2 R2000FastSwitchInventoryEnum = 0x01
	ANTENNA3 R2000FastSwitchInventoryEnum = 0x02
	ANTENNA4 R2000FastSwitchInventoryEnum = 0x03
	DISABLED R2000FastSwitchInventoryEnum = 0xFF
)

var FastSwitchInventory = struct {
	ANTENNA1 []byte
	ANTENNA2 []byte
	ANTENNA3 []byte
	ANTENNA4 []byte
	DISABLED []byte
}{
	ANTENNA1: []byte{byte(ANTENNA1)},
	ANTENNA2: []byte{byte(ANTENNA2)},
	ANTENNA3: []byte{byte(ANTENNA3)},
	ANTENNA4: []byte{byte(ANTENNA4)},
	DISABLED: []byte{byte(DISABLED)},
}
var FastSwitchInventoryMap = map[string]R2000FastSwitchInventoryEnum{
	"ANTENNA1": ANTENNA1,
	"ANTENNA2": ANTENNA2,
	"ANTENNA3": ANTENNA3,
	"ANTENNA4": ANTENNA4,
	"DISABLED": DISABLED,
}
