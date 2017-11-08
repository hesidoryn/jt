package server

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/hesidoryn/jt/storage"
)

var handlers = map[string]func(args [][]byte, w *bufio.Writer){}

const (
	resultOK   = "+OK"
	resultPONG = "+PONG"

	errorUnknownCommand = "-ERR unknown command '%s'"
	errorWrongArguments = "-ERR wrong number of arguments for '%s' command"
	errorSyntaxError    = "-ERR syntax error"
	errorNoSuchKey      = "-ERR no such key"
	errorIsNotInteger   = "-ERR value is not an integer or out of range"
	errorIsNotFloat     = "-ERR dict value is not a float"
	errorNotFound       = "-ERR"
	errorWrongType      = "-WRONGTYPE Operation against a key holding the wrong kind of value"
	errorProtocolError  = "-ERR Protocol error: unbalanced quotes in request"
)

// Init function inits tcp server
func Init(port string) {
	storage.GetStorage()
	initHandlers()

	listen, err := net.Listen("tcp4", ":"+port)
	defer listen.Close()
	if err != nil {
		log.Fatalf("Socket listen port %s failed,%s", port, err)
		os.Exit(1)
	}
	log.Printf("Begin listen port: %s", port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	w := bufio.NewWriter(conn)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Bytes()
		args, err := parseArgs(ln)
		if err != nil {
			sendResult(errorProtocolError, w)
			return
		}
		if len(args) == 0 {
			continue
		}

		command := string(bytes.ToUpper(args[0]))
		handler, ok := handlers[command]
		if !ok {
			sendResult(fmt.Sprintf(errorUnknownCommand, args[0]), w)
			continue
		}

		handler(args, w)
	}
}

func initHandlers() {
	initCommonHandlers()
	initStringHandlers()
	initListHandlers()
	initDictHandlers()
}

// if len(fs) < 2 {
// 	w.WriteString("This is an in-memory database \n" +
// 		"Use SET, GET, DEL like this: \n" +
// 		"SET key value \n" +
// 		"GET key \n" +
// 		"DEL key \n\n" +
// 		"For example - try these commands: \n" +
// 		"SET fav chocolate \n" +
// 		"GET fav \n\n")
// 	continue
// }
