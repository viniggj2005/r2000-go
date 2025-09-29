package dtos

import "github.com/viniggj2005/r2000-go/enums"

type BuildFrame struct {
	Command enums.R2000CommandsEnum
	Params  []byte
}
