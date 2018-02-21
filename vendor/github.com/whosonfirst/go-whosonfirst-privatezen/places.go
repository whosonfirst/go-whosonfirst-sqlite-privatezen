package privatezen

import (
	"time"
)

type Place struct {
	Id      *WOFPlace
	Body    string
	Created *time.Time
}
