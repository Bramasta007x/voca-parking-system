package parking

import "time"

type Car struct {
	RegistrationNumber string
	Color              string
	SlotNumber         int
	EntryTime          time.Time
}
