package parking

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type ParkingLot struct {
	capacity  int            // Kapasitas maksimum slot
	slots     map[int]*Car   // Map slot number ke objek Car
	carToSlot map[string]int // Map nomor registrasi ke slot
}

func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		capacity:  capacity,
		slots:     make(map[int]*Car),
		carToSlot: make(map[string]int),
	}
}

func (pl *ParkingLot) Park(car *Car) (int, error) {
	// Check if car already parked
	if _, exsits := pl.carToSlot[car.RegistrationNumber]; exsits {
		return 0, fmt.Errorf("car %s is already parked", car.RegistrationNumber)
	}

	// Find nearest available slot
	for i := 1; i <= pl.capacity; i++ {
		if _, occupied := pl.slots[i]; !occupied {
			car.SlotNumber = i
			car.EntryTime = time.Now()
			pl.slots[i] = car
			pl.carToSlot[car.RegistrationNumber] = i
			return i, nil
		}
	}

	return 0, fmt.Errorf("Sorry, parking lot is full")
}

func (pl *ParkingLot) LeaveByRegistration(registrationNumber string, hours int) (int, int, error) {
	slotNumber, exists := pl.carToSlot[registrationNumber]
	if !exists {
		return 0, 0, fmt.Errorf("Registration number %s not found", registrationNumber)
	}

	// Calculate fee: $10 for first 2 hours, $10 for each additional hour
	fee := 10
	if hours > 2 {
		fee += (hours - 2) * 10
	}

	delete(pl.slots, slotNumber)
	delete(pl.carToSlot, registrationNumber)

	return slotNumber, fee, nil
}

func (pl *ParkingLot) Status() string {
	if len(pl.slots) == 0 {
		return "Parking lot is empty"
	}

	var result strings.Builder
	result.WriteString("Slot No. Registration No.\n")

	// Get sorted slot numbers
	var slotNumbers []int
	for slotNum := range pl.slots {
		slotNumbers = append(slotNumbers, slotNum)
	}
	sort.Ints(slotNumbers)

	for _, slotNum := range slotNumbers {
		car := pl.slots[slotNum]
		result.WriteString(fmt.Sprintf("%d %s\n", slotNum, car.RegistrationNumber))
	}

	return strings.TrimSpace(result.String())
}
