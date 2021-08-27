package internal

import (
	"regexp"
)

type UserCommand struct {
	Command   string
	Arguments []string
}

func ArgumentFilter(commandList []string) UserCommand {
	regularExpressionValidator := regexp.MustCompile("(?m)-")
	commandSet := false

	var commandDefinition UserCommand

	for _, argument := range commandList {
		isMatch := regularExpressionValidator.MatchString(argument)
		if !isMatch && !commandSet {
			commandDefinition.Command = argument
			commandSet = true
		} else if isMatch {
			commandDefinition.Arguments = append(commandDefinition.Arguments, argument)
		}
	}

	return commandDefinition
}
