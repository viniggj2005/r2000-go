package enums

type R2000MemoryBankEnum byte

const (
	EPC      R2000MemoryBankEnum = 0x01
	TID      R2000MemoryBankEnum = 0x02
	USERB    R2000MemoryBankEnum = 0x03
	RESERVED R2000MemoryBankEnum = 0x00
)

var MemoryBank = struct {
	EPC      []byte
	TID      []byte
	USERB    []byte
	RESERVED []byte
}{
	EPC:      []byte{byte(EPC)},
	TID:      []byte{byte(TID)},
	USERB:    []byte{byte(USERB)},
	RESERVED: []byte{byte(RESERVED)},
}
