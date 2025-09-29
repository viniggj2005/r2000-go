package enums

type R2000StateEnum byte

const (
	OFF R2000StateEnum = 0x00
	ON  R2000StateEnum = 0x01
)

var State = struct {
	ON  []byte
	OFF []byte
}{
	ON:  []byte{byte(ON)},
	OFF: []byte{byte(OFF)},
}
