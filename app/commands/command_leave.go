package commands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/amogmish/parkingLot/app/models/parking"
)

type CommandLeave struct {
	Command
	SlotNumber int
}

func (cl *CommandLeave) ParseArguments(args string) error {
	cl.Arguments = strings.Split(args, " ")
	if !cl.ValidateInput() {
		return errInvalidInput
	}
	value, err := strconv.ParseInt(args, 0, 64)
	if err != nil {
		return errInvalidInput
	}
	cl.SlotNumber = int(value)
	return nil
}

func (cl *CommandLeave) Clear() {
	cl.Arguments = nil
	cl.SlotNumber = 0
}

func (cl *CommandLeave) ValidateInput() bool {
	return len(cl.Arguments) == 1
}

func (cl *CommandLeave) Run() (string, error) {
	var output string
	park := parking.Get()
	slotNumber := cl.SlotNumber
	if err := park.RemoveCarFromSlot(slotNumber); err != nil {
		return output, err
	}
	output = fmt.Sprintf("Slot number %d is free", slotNumber)
	return output, nil
}
