package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net"

	"github.com/radenrishwan/cocytus"
)

var (
	PORT = flag.String("port", "6379", "port to listen on (default 6379)")
)

var (
	STR   = "+"
	ERR   = "-"
	INT   = ":"
	BULK  = "$"
	ARRAY = "*"
)

var (
	GET = "GET"
	SET = "SET"
)

var (
	CRLF = "\r\n"
)

func main() {
	server, err := net.Listen("tcp", ":"+*PORT)
	if err != nil {
		log.Fatalln(err)
	}

	defer server.Close()
	log.Println("Server is running on port " + *PORT)

	for {
		conn, err := server.Accept()
		if err != nil {
			slog.Error("Error accepting connection", "ERR", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn io.ReadWriteCloser) {
	defer conn.Close()

	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)

	for {
		cmd := cocytus.NewCommand()
		err := cmd.Parse(r)
		if err != nil {
			if err == io.EOF {
				// client closed connection
				slog.Info("EOF Receive", "ERR", "maybe client closed connection")
				return
			}

			slog.Error("Error reading command", "ERR", err)

			// write error here
			writeError(w, err.Error())
			w.Flush()
		}

		fmt.Println(cmd)

		w.WriteString(STR + "PONG" + CRLF)

		// HANDLE COMMAND HERE
		w.Flush()
	}
}

func writeError(w *bufio.Writer, msg string) {
	w.WriteString(ERR + msg + CRLF)
	w.Flush()
}
