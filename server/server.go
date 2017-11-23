package server

const (
	cmdAuth = "AUTH"
	cmdPing = "PING"
	cmdSave = "SAVE"
)

func auth(c jtContext) {
	if password == "" {
		c.sendResult(errorNoPassword)
		return
	}

	c.client.password = string(c.args[1])
	if c.client.password != password {
		c.client.isAuthorized = false
		c.sendResult(errorInvalidPassword)
		return
	}

	c.client.isAuthorized = true
	c.sendResult(resultOK)
}

func ping(c jtContext) {
	c.sendResult(resultPONG)
}

func save(c jtContext) {
	jtStorage.Save()
	c.sendResult(resultOK)
}
