package enums

type R2000BeeperEnum byte

const (
	QuietEnum               R2000BeeperEnum = 0x00
	AfterInventoryRoundEnum R2000BeeperEnum = 0x01
	AfterEveryTagEnum       R2000BeeperEnum = 0x02
)

var BeeperMode = struct {
	Quiet               []byte
	AfterInventoryRound []byte
	AfterEveryTag       []byte
}{
	Quiet:               []byte{byte(QuietEnum)},
	AfterInventoryRound: []byte{byte(AfterInventoryRoundEnum)},
	AfterEveryTag:       []byte{byte(AfterEveryTagEnum)},
}
