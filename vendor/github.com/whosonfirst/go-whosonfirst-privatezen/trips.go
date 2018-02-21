package privatezen

import (
	"time"
)

var states map[int64]string

func init() {

	states = map[int64]string{
		0: "unknown",
		1: "tentative",
		2: "confirmed",
		3: "i want this to happen",
	}
}

func IsValidTripStatus(s string) bool {

	valid := false

	for _, label := range states {

		if s == label {
			valid = true
			break
		}
	}

	return valid
}

func IsValidTripStatusId(id int64) bool {

	_, ok := states[id]
	return ok
}

type TripStatus struct {
	Id    int64
	Label string
}

type Trip struct {
	Id          int64
	Name        string
	Arrival     *time.Time
	Departure   *time.Time
	Destination *WOFPlace // wof_id
	Status      *TripStatus
	Notes       string
}
