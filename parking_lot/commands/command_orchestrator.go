package commands

import (
	"fmt"
	"strings"

	"github.com/amogmish/parkingLot/parking_lot/models/parking"
)

// ICommand is interface all registered command
type ICommand interface {
	ParseArguments(string) error
	Clear()
	ValidateInput() bool
	Run() (string, error)
}

// Command is object struct for manage command from user and control application
type Command struct {
	Arguments   []string
	CommandList map[string]ICommand
}

var (
	errInvalidInput        = fmt.Errorf("invalid parameter(s), please provide valid input")
	errParkingNotAvailable = fmt.Errorf("please create_parking_lot first")
)

// InitCommandOrchestrator is for registering all command to control application
func InitCommandOrchestrator() *Command {
	cmd := new(Command)
	cmd.CommandList = make(map[string]ICommand)
	cmd.CommandList["create_parking_lot"] = new(CommandCreateParkingLot)
	cmd.CommandList["park"] = new(CommandPark)
	cmd.CommandList["leave"] = new(CommandLeave)
	cmd.CommandList["status"] = new(CommandStatus)
	return cmd
}

// Run command with parameter string entered by user
func (cmd *Command) Run(command string) string {
	cmds := strings.SplitN(command, " ", 2)
	menu := cmds[0]
	cmdChild, ok := cmd.CommandList[menu]
	if cmdChild == nil || !ok {
		return fmt.Sprintf("%s: command not found\n", menu)
	}

	// clearing command argument(s)
	cmdChild.Clear()

	// capture command argument(s)
	if len(cmds) > 1 {
		cmdChild.ParseArguments(cmds[1])
	}

	valid := cmdChild.ValidateInput()
	if !valid {
		return fmt.Sprintf("%s: invalid input provided, please provide valid input\n", menu)
	}

	park := parking.Get()
	if park == nil && menu != "create_parking_lot" {
		return fmt.Sprintf("%v\n", errParkingNotAvailable)
	}

	output, err := cmdChild.Run()
	if err != nil {
		return fmt.Sprintf("%v\n", err)
	}
	return output
}
