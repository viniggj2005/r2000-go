package enums

type R2000MemoryBankEnum byte

const (
	RESERVED R2000MemoryBankEnum = 0x00
	EPC      R2000MemoryBankEnum = 0x01
	TID      R2000MemoryBankEnum = 0x02
	USERB    R2000MemoryBankEnum = 0x03
)
