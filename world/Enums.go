package world

type ReturnValue int
const (
	RET_NOERROR ReturnValue = iota
	RET_NOTPOSSIBLE
	RET_TILEPOINT_NOTFOUND
	RET_TILEPOINTLAYER_NOTFOUND
	RET_TELEPORTED
)