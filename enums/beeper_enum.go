package enums

type R2000BeeperEnum byte

const (
	QUIET                 R2000BeeperEnum = 0x00 // sem som
	AFTER_INVENTORY_ROUND R2000BeeperEnum = 0x01 // beep após cada rodada de inventário
	AFTER_EVERY_TAG       R2000BeeperEnum = 0x02 // beep após cada tag lida
)
