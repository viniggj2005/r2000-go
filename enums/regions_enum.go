package enums

type R2000RegionsEnum byte

const (
	FCC  R2000RegionsEnum = 1
	ETSI R2000RegionsEnum = 2
	CHN  R2000RegionsEnum = 3
	USER R2000RegionsEnum = 4
)

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
