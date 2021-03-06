// Package server implements tcp server to handle clients requests.
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

type JTServer struct {
	routes   map[string]jtHandlerFunc
	password string
}

func (s *JTServer) Handle(command string, handler jtHandlerFunc, ms ...jtMiddlewareFunc) {
	result := handler
	for _, m := range ms {
		result = m(result)
	}
	s.routes[command] = result
}

type client struct {
	conn         net.Conn
	w            *bufio.Writer
	sc           *bufio.Scanner
	password     string
	isAuthorized bool
}

var (
	password  string
	jtStorage *storage.JTStorage
)

const (
	resultOK             = "+OK"
	resultPONG           = "+PONG"
	resultDefaultString  = "$-1"
	resultDefaultInteger = "-1"
	resultDefaultList    = "*0"
	resultDefaultDict    = "*1\r\n$-1"

	errorUnknownCommand  = "-ERR unknown command '%s'"
	errorWrongArguments  = "-ERR wrong number of arguments for '%s' command"
	errorSyntaxError     = "-ERR syntax error"
	errorNoSuchKey       = "-ERR no such key"
	errorIsNotInteger    = "-ERR value is not an integer or out of range"
	errorIsNotFloat      = "-ERR dict value is not a float"
	errorNoPassword      = "-ERR Client sent AUTH, but no password is set"
	errorWrongType       = "-WRONGTYPE Operation against a key holding the wrong kind of value"
	errorInvalidPassword = "-ERR invalid password"
	errorNoAuth          = "-NOAUTH Authentication required."
	errorProtocolError   = "-ERR Protocol error: unbalanced quotes in request"
)

// Init function inits tcp server
func Init(config config.Config) {
	jtStorage = storage.Init(config.DumpFile)

	password = config.Password
	jtServer := &JTServer{
		routes: map[string]jtHandlerFunc{},
	}
	jtServer.loadRoutes()

	listen, err := net.Listen("tcp", ":"+config.Port)
	defer listen.Close()
	if err != nil {
		log.Fatalf("Socket listen port %s failed,%s", config.Port, err)
		os.Exit(1)
	}
	log.Printf("JT begins listen port: %s", config.Port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		isAuth := false
		if password == "" {
			isAuth = true
		}
		c := &client{
			conn:         conn,
			w:            bufio.NewWriter(conn),
			sc:           bufio.NewScanner(conn),
			isAuthorized: isAuth,
		}
		go jtServer.handleConnection(c)
	}
}

func (s *JTServer) handleConnection(c *client) {
	defer c.conn.Close()

	context := jtContext{client: c}
	for c.sc.Scan() {
		ln := c.sc.Bytes()
		if len(ln) == 0 {
			continue
		}

		args, err := parseArgs(ln)
		if err != nil {
			context.sendResult(errorProtocolError)
			return
		}

		command := string(bytes.ToUpper(args[0]))
		handler, ok := s.routes[command]
		if !ok {
			context.sendResult(fmt.Sprintf(errorUnknownCommand, command))
			continue
		}

		context.command, context.args = command, args
		handler(context)
	}
}
