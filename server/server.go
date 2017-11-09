package server

import (
	"fmt"
)

func initServerHandlers() {
	handlers["AUTH"] = handlerAuth
	handlers["PING"] = handlerPing
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
