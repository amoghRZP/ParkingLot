package commands

import (
	"fmt"
	"strings"

	"github.com/amogmish/parkingLot/app/models/parking"
)

type CommandSlotNumberCarNumber struct {
	Command
	CarNumber string
}

func (csncn *CommandSlotNumberCarNumber) ParseArguments(args string) error {
	csncn.Arguments = strings.Split(args, " ")
	if !csncn.ValidateInput() {
		return errInvalidInput
	}
	csncn.CarNumber = csncn.Arguments[0]
	return nil
}

func (csncn *CommandSlotNumberCarNumber) Clear() {
	csncn.Arguments = nil
	csncn.CarNumber = ""
}

func (csncn *CommandSlotNumberCarNumber) ValidateInput() bool {
	return len(csncn.Arguments) == 1 && csncn.Arguments[0] != ""
}

func (csncn *CommandSlotNumberCarNumber) Run() (string, error) {
	var output string
	slot := parking.Get().GetSlotByCarNumber(csncn.CarNumber)
	if slot == nil {
		return "Not found", nil
	}
	output = fmt.Sprint(slot.Number)
	return output, nil
}
