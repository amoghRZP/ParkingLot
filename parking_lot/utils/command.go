package utils

func RunCommand(command string) string {
	var result string
	switch command {
	case "park":
		result = park(command)
	case "leave":
		result = leave(command)
	case "status":
		result = status(command)
	case "create_parking_lot":
		result = createParkingLot(command)
	case "slot_numbers_for_cars_with_colour":
		result = slotNumberCarNumber(command)
	case "slot_number_for_registration_number":
		result = slotNumberCarColor(command)
	case "registration_numbers_for_cars_with_colour":
		result = registrationNumberWithColor(command)
	}
	return result
}

func park(command string) string {
	return ""
}

func leave(command string) string {
	return ""
}

func status(command string) string {
	return ""
}

func createParkingLot(command string) string {
	return ""
}

func slotNumberCarNumber(command string) string {
	return ""
}

func slotNumberCarColor(command string) string {
	return ""
}

func registrationNumberWithColor(command string) string {
	return ""
}
