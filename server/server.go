package server

import (
	"fmt"

	"github.com/hesidoryn/jt/storage"
)

const (
	cmdAuth = "AUTH"
	cmdPing = "PING"
	cmdSave = "SAVE"
)

func initServerHandlers() {
	handlers[cmdAuth] = handlerAuth
	handlers[cmdPing] = handlerPing
	handlers[cmdSave] = handlerSave
}

func handlerAuth(args [][]byte, c *client) {
	if len(args) != 2 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	if password == "" {
		sendResult(errorNoPassword, c.w)
		return
	}

	c.password = string(args[1])
	if c.password != password {
		c.isAuthorized = false
		sendResult(errorInvalidPassword, c.w)
		return
	}

	c.isAuthorized = true
	sendResult(resultOK, c.w)
}

func handlerPing(args [][]byte, c *client) {
	if len(args) != 1 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	sendResult(resultPONG, c.w)
}

func handlerSave(args [][]byte, c *client) {
	if len(args) != 1 {
		sendResult(fmt.Sprintf(errorWrongArguments, args[0]), c.w)
		return
	}

	storage.Save()
	sendResult(resultOK, c.w)
}
