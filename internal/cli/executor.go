package cli

import (
	"fmt"
	"strconv"
	"strings"

	"parking-system/internal/parking"
)

var parkingLot *parking.ParkingLot

func ExecuteCommand(command string) string {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return "Invalid command"
	}

	switch parts[0] {
	case "create_parking_lot":
		return createParkingLot(parts)
	case "park":
		return parkCar(parts)
	case "leave":
		return leaveCar(parts)
	case "status":
		return getStatus()
	default:
		return "Unknown command"
	}
}

func createParkingLot(parts []string) string {
	if len(parts) != 2 {
		return "Invalid command format. Usage: create_parking_lot <capacity>"
	}

	capacity, err := strconv.Atoi(parts[1])
	if err != nil {
		return "Invalid capacity number"
	}

	parkingLot = parking.NewParkingLot(capacity)
	return fmt.Sprintf("Created a parking lot with %d slots", capacity)
}

func parkCar(parts []string) string {
	if len(parts) != 2 {
		return "Invalid command format. Usage: park <registration_number>"
	}

	if parkingLot == nil {
		return "Parking lot not created"
	}

	car := &parking.Car{
		RegistrationNumber: parts[1],
		Color:              "", // Default empty color
	}

	slotNumber, err := parkingLot.Park(car)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("Allocated slot number: %d", slotNumber)
}

func leaveCar(parts []string) string {
	if len(parts) != 3 {
		return "Invalid command format. Usage: leave <registration_number> <hours>"
	}

	if parkingLot == nil {
		return "Parking lot not created"
	}

	hours, err := strconv.Atoi(parts[2])
	if err != nil {
		return "Invalid hours"
	}

	slotNumber, fee, err := parkingLot.LeaveByRegistration(parts[1], hours)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("Registration number %s with Slot Number %d is free with Charge $%d", parts[1], slotNumber, fee)
}

func getStatus() string {
	if parkingLot == nil {
		return "Parking lot not created"
	}

	return parkingLot.Status()
}
