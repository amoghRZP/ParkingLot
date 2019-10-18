package commands

import (
	"fmt"
	"strings"

	"github.com/amogmish/parkingLot/parking_lot/models/parking"
)

type CommandStatus struct {
	Command
}

func (cs *CommandStatus) ParseArguments(args string) error {
	return nil
}

func (cs *CommandStatus) Clear() {
	cs.Arguments = nil
}

func (cs *CommandStatus) ValidateInput() bool {
	return true
}

func (cs *CommandStatus) Run() (string, error) {
	var list = []string{fmt.Sprintf("%-12s%-20s%-10s", "Slot No.", "Registration No", "Colour")}
	filledSlots := parking.Get().GetOccupiedSlots()
	for _, sl := range filledSlots {
		cr := sl.Car
		list = append(list, fmt.Sprintf("%-12v%-20v%-10v", sl.Number, cr.Number, cr.Color))
	}
	output := strings.Join(list, "\n")
	return output, nil
}
