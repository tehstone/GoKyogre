package framework

import (
	"fmt"
)

type (
	Command func(Context)

	CmdMap map[string]Command

	CommandHandler struct {
		cmds CmdMap
	}
)

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{make(CmdMap)}
}

func (handler CommandHandler) GetCmds() CmdMap {
	return handler.cmds
}

func (handler CommandHandler) Get(name string) (*Command, bool) {
	cmd, found := handler.cmds[name]
	return &cmd, found
}

func (handler CommandHandler) Register(name string, command Command) {
	handler.cmds[name] = command
	if len(name) > 1 {
		handler.cmds[name] = command
		fmt.Println("Registered command,", name)
	} else {
		fmt.Println("Command not registered.")
	}
}
