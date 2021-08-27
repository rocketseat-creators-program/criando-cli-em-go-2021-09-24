package main

import (
	"fmt"
	"os"

	"github.com/GuilhermeCaruso/expertscli/commands"
	"github.com/GuilhermeCaruso/expertscli/internal"
)

var commandList = []internal.Command{
	new(commands.Start),
}

func main() {
	if err := internal.CommandInit("expertscli").Start(commandList); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
