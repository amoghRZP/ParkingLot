package commands

import (
	"fmt"
	"strings"

	"github.com/amogmish/parkingLot/app/models/parking"
)

type CommandSlotNumberCarColor struct {
	Command
	CarColor string
}

func (csncc *CommandSlotNumberCarColor) ParseArguments(args string) error {
	csncc.Arguments = strings.Split(args, " ")
	if !csncc.ValidateInput() {
		return errInvalidInput
	}
	csncc.CarColor = csncc.Arguments[0]
	return nil
}

func (csncc *CommandSlotNumberCarColor) Clear() {
	csncc.Arguments = nil
	csncc.CarColor = ""
}

func (csncc *CommandSlotNumberCarColor) ValidateInput() bool {
	return len(csncc.Arguments) == 1 && csncc.Arguments[0] != ""
}

func (this *CommandSlotNumberCarColor) Run() (string, error) {
	var output string
	var list []string
	slots := parking.Get().GetSlotsByCarColor(this.CarColor)
	if slots == nil {
		return "Not found", nil
	}
	for _, sl := range slots {
		list = append(list, fmt.Sprint(sl.Number))
	}
	output = strings.Join(list, ", ")
	return output, nil
}
