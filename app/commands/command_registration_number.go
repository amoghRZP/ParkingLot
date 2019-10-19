package commands

import (
	"strings"

	"github.com/amogmish/parkingLot/app/models/parking"
)

type CommandRegistrationNumber struct {
	Command
	CarColor string
}

func (crn *CommandRegistrationNumber) ParseArgs(args string) error {
	crn.Arguments = strings.Split(args, " ")
	if !crn.ValidateInput() {
		return errInvalidInput
	}
	crn.CarColor = crn.Arguments[0]
	return nil
}

func (crn *CommandRegistrationNumber) Clear() {
	crn.Arguments = nil
	crn.CarColor = ""
}

func (crn *CommandRegistrationNumber) ValidateInput() bool {
	return len(crn.Arguments) == 1 && crn.Arguments[0] != ""
}

func (this *CommandRegistrationNumber) Run() (string, error) {
	var output string
	var list []string
	slots := parking.Get().GetSlotsByCarColor(this.CarColor)
	if slots == nil {
		return "Not found", nil
	}
	for _, sl := range slots {
		list = append(list, sl.GetCarNumberInSlot())
	}
	output = strings.Join(list, ", ")
	return output, nil
}
