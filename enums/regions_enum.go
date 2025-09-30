package enums

type R2000RegionsEnum byte

const (
	FCC  R2000RegionsEnum = 1
	ETSI R2000RegionsEnum = 2
	CHN  R2000RegionsEnum = 3
	USER R2000RegionsEnum = 4
)

var Regions = struct {
	FCC  byte
	CHN  byte
	ETSI byte
	USER byte
}{
	FCC:  byte(FCC),
	CHN:  byte(CHN),
	ETSI: byte(ETSI),
	USER: byte(USER),
}

func (r R2000RegionsEnum) String() string {
	switch r {
	case ETSI:
		return "ETSI"
	case FCC:
		return "FCC"
	case CHN:
		return "CHN"
	case USER:
		return "USER"
	default:
		return "UNKNOWN"
	}
}

var RegionsMap = map[string]R2000RegionsEnum{
	"FCC":  FCC,
	"ETSI": ETSI,
	"CHN":  CHN,
	"USER": USER,
}
