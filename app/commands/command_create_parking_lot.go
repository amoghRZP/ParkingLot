package commands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/amogmish/parkingLot/app/models/parking"
)

type CommandCreateParkingLot struct {
	Command
	CreationCapacity int
}

func (cpl *CommandCreateParkingLot) ParseArguments(args string) error {
	cpl.Arguments = strings.Split(args, " ")
	if !cpl.ValidateInput() {
		return errInvalidInput
	}
	value, err := strconv.ParseUint(args, 0, 64)
	if err != nil {
		return errInvalidInput
	}
	cpl.CreationCapacity = int(value)
	return nil
}

func (cpl *CommandCreateParkingLot) Clear() {
	cpl.Arguments = nil
	cpl.CreationCapacity = 0
}

func (cpl *CommandCreateParkingLot) ValidateInput() bool {
	return len(cpl.Arguments) == 1
}

func (cpl *CommandCreateParkingLot) Run() (string, error) {
	var output string
	obj := parking.New(cpl.CreationCapacity)
	if obj == nil {
		return output, fmt.Errorf("Error")
	}
	output = fmt.Sprintf("Created a parking lot with %d slots", cpl.CreationCapacity)
	return output, nil
}
