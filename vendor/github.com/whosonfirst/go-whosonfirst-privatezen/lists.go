package privatezen

import (
	"time"
)

type List struct {
	Id      int64
	Name    string
	Created *time.Time
	Items   []*ListItem
}

type ListItem struct {
	List  *List
	Foo   *WOFPlace
	Index int // position
}
