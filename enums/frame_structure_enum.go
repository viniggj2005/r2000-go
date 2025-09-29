package enums

type FrameStructureEnum byte

const (
	HEADER     FrameStructureEnum = 0xA0
	MIN_LENGTH FrameStructureEnum = 5
)

var FrameStructure = struct {
	HEADER     []byte
	MIN_LENGTH []byte
}{
	HEADER:     []byte{byte(HEADER)},
	MIN_LENGTH: []byte{byte(MIN_LENGTH)},
}
