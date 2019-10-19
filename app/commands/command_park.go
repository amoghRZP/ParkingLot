package commands

import (
	"fmt"
	"strings"

	"github.com/amogmish/parkingLot/app/models/car"
	"github.com/amogmish/parkingLot/app/models/parking"
)

type CommandPark struct {
	Command
	Car *car.Car
}

func (cp *CommandPark) ParseArguments(args string) error {
	cp.Arguments = strings.Split(args, " ")
	if !cp.ValidateInput() {
		return errInvalidInput
	}
	cp.Car = car.New(cp.Arguments[0], cp.Arguments[1])
	return nil
}

func (cp *CommandPark) Clear() {
	cp.Arguments = nil
	cp.Car = nil
}

func (cp *CommandPark) ValidateInput() bool {
	return len(cp.Arguments) == 2
}

func (cp *CommandPark) Run() (string, error) {
	var output string
	sl, err := parking.Get().AddCar(*cp.Car)
	if err != nil {
		return output, err
	}
	output = fmt.Sprintf("Occupied slot number: %d", sl.Number)
	return output, err
}
