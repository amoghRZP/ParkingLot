package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/amogmish/parkingLot/parking_lot/utils"
)

func main() {

	// check if file input exists and read from file
	if len(os.Args) > 1 && os.Args[1] != "" {
		cmdLineFile, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer cmdLineFile.Close()
		cmdScanner := bufio.NewScanner(cmdLineFile)
		for cmdScanner.Scan() {
			cmdInput := cmdScanner.Text()
			cmdInput = strings.TrimRight(cmdInput, "\n")
			if cmdInput != "" {
				output := utils.RunCommand(cmdInput)
				fmt.Println(output)
			}
		}

		if err := cmdScanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		cmdInput, _ := reader.ReadString('\n')
		cmdInput = strings.TrimRight(cmdInput, "\n")
		if cmdInput != "" {
			output := utils.RunCommand(cmdInput)
			fmt.Println(output)
		}
	}
}
