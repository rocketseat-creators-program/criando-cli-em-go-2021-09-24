package main

import (
	"fmt"

	"github.com/GuilhermeCaruso/expertscli/commands/message"
	"github.com/GuilhermeCaruso/expertscli/commands/start"

	"github.com/GuilhermeCaruso/expertscli/internal"
)

var commandList = []internal.Command{
	new(start.Start),
	new(message.Message),
}

func main() {
	if err := internal.CommandInit("expertscli").Start(commandList); err != nil {
		fmt.Println(err.Error())
	}
}
