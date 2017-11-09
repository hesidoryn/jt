package server

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/hesidoryn/jt/config"
	"github.com/hesidoryn/jt/storage"
)

type client struct {
	conn         net.Conn
	w            *bufio.Writer
	sc           *bufio.Scanner
	password     string
	isAuthorized bool
}

var (
	handlers = map[string]func(args [][]byte, c *client){}

	password = ""
)

const (
	resultOK   = "+OK"
	resultPONG = "+PONG"

	errorUnknownCommand  = "-ERR unknown command '%s'"
	errorWrongArguments  = "-ERR wrong number of arguments for '%s' command"
	errorSyntaxError     = "-ERR syntax error"
	errorNoSuchKey       = "-ERR no such key"
	errorIsNotInteger    = "-ERR value is not an integer or out of range"
	errorIsNotFloat      = "-ERR dict value is not a float"
	errorNoPassword      = "-ERR Client sent AUTH, but no password is set"
	errorNotFound        = "-ERR"
	errorWrongType       = "-WRONGTYPE Operation against a key holding the wrong kind of value"
	errorInvalidPassword = "-ERR invalid password"
	errorNoAuth          = "-NOAUTH Authentication required."
	errorProtocolError   = "-ERR Protocol error: unbalanced quotes in request"
)

// Init function inits tcp server
func Init(config config.Config) {
	storage.GetStorage()
	initHandlers()
	password = config.Password

	listen, err := net.Listen("tcp4", ":"+config.Port)
	defer listen.Close()
	if err != nil {
		log.Fatalf("Socket listen port %s failed,%s", config.Port, err)
		os.Exit(1)
	}
	log.Printf("Begin listen port: %s", config.Port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		writer := bufio.NewWriter(conn)
		scanner := bufio.NewScanner(conn)
		isAuth := false
		if password == "" {
			isAuth = true
		}
		c := &client{
			conn:         conn,
			w:            writer,
			sc:           scanner,
			isAuthorized: isAuth,
		}
		go handleConnection(c)
	}
}

func handleConnection(c *client) {
	defer c.conn.Close()

	scanner := bufio.NewScanner(c.conn)
	for scanner.Scan() {
		ln := scanner.Bytes()
		args, err := parseArgs(ln)
		if err != nil {
			sendResult(errorProtocolError, c.w)
			return
		}
		if len(args) == 0 {
			continue
		}

		command := string(bytes.ToUpper(args[0]))
		handler, ok := handlers[command]
		if !ok {
			sendResult(fmt.Sprintf(errorUnknownCommand, args[0]), c.w)
			continue
		}

		if command != "AUTH" && !c.isAuthorized {
			sendResult(errorNoAuth, c.w)
			continue
		}
		handler(args, c)
	}
}

func initHandlers() {
	initServerHandlers()
	initKeysHandlers()
	initStringHandlers()
	initListHandlers()
	initDictHandlers()
}
