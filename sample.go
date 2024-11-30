package cocytus

// import (
// 	"bufio"
// 	"flag"
// 	"fmt"
// 	"io"
// 	"log"
// 	"log/slog"
// 	"net"
// 	"strings"
// )

// var (
// 	PORT = flag.String("port", "6379", "port to listen on (default 6379)")
// )

// var (
// 	STR   = "+"
// 	ERR   = "-"
// 	INT   = ":"
// 	BULK  = "$"
// 	ARRAY = "*"
// )

// var (
// 	GET = "GET"
// 	SET = "SET"
// )

// var (
// 	CRLF = "\r\n"
// )

// func main() {
// 	server, err := net.Listen("tcp", ":"+*PORT)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	defer server.Close()
// 	log.Println("Server is running on port " + *PORT)

// 	for {
// 		conn, err := server.Accept()
// 		if err != nil {
// 			slog.Error("Error accepting connection", "ERR", err)
// 		}

// 		go handleConnection(conn)
// 	}
// }

// func handleConnection(conn io.ReadWriteCloser) {
// 	defer conn.Close()

// 	r := bufio.NewReader(conn)
// 	w := bufio.NewWriter(conn)

// 	for {
// 		firstByte, err := r.ReadByte()
// 		if err != nil {
// 			if err == io.EOF {
// 				slog.Error("Eror reading command", "ERR", "EOF, maybe client closed connection")
// 				return
// 			}

// 			slog.Error("Error reading command", "ERR", err)
// 			continue
// 		}

// 		switch string(firstByte) {
// 		case ARRAY:
// 			// Read array length
// 			length, err := readLength(r)
// 			if err != nil {
// 				writeError(w, "Invalid array length")
// 				continue
// 			}

// 			// Read array elements
// 			command := make([]string, length)
// 			for i := 0; i < length; i++ {
// 				// Each array element should be a bulk string
// 				bulkByte, err := r.ReadByte()
// 				if err != nil || string(bulkByte) != BULK {
// 					writeError(w, "Invalid command format")
// 					continue
// 				}

// 				// Read string length
// 				strLen, err := readLength(r)
// 				if err != nil {
// 					writeError(w, "Invalid string length")
// 					continue
// 				}

// 				// Read string content
// 				str := make([]byte, strLen)
// 				if _, err := io.ReadFull(r, str); err != nil {
// 					writeError(w, "Error reading string")
// 					continue
// 				}

// 				// Read CRLF
// 				if _, err := r.ReadString('\n'); err != nil {
// 					writeError(w, "Expected CRLF")
// 					continue
// 				}

// 				command[i] = string(str)
// 			}

// 			// Handle command
// 			handleCommand(command, w)
// 			w.Flush()

// 		default:
// 			writeError(w, "Unknown command type")
// 			w.Flush()
// 		}
// 	}
// }

// func readLength(r *bufio.Reader) (int, error) {
// 	line, err := r.ReadString('\n')
// 	if err != nil {
// 		return 0, err
// 	}

// 	// Remove CRLF
// 	line = line[:len(line)-2]

// 	// Parse length
// 	length := 0
// 	for _, ch := range line {
// 		if ch < '0' || ch > '9' {
// 			return 0, fmt.Errorf("invalid length")
// 		}
// 		length = length*10 + int(ch-'0')
// 	}
// 	return length, nil
// }

// func writeError(w *bufio.Writer, msg string) {
// 	w.WriteString(ERR + msg + CRLF)
// 	w.Flush()
// }

// func handleCommand(cmd []string, w *bufio.Writer) {
// 	if len(cmd) == 0 {
// 		writeError(w, "Empty command")
// 		return
// 	}

// 	switch strings.ToUpper(cmd[0]) {
// 	case "PING":
// 		w.WriteString(STR + "PONG" + CRLF)

// 	case GET:
// 		if len(cmd) != 2 {
// 			writeError(w, "Wrong number of arguments for GET command")
// 			return
// 		}
// 		// TODO: Implement GET logic
// 		w.WriteString(STR + "value" + CRLF)

// 	case SET:
// 		if len(cmd) != 3 {
// 			writeError(w, "Wrong number of arguments for SET command")
// 			return
// 		}
// 		// TODO: Implement SET logic
// 		w.WriteString(STR + "OK" + CRLF)

// 	default:
// 		writeError(w, "Unknown command '"+cmd[0]+"'")
// 	}
// }
