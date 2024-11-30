package cocytus

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

var (
	STR   = "+"
	ERR   = "-"
	INT   = ":"
	BULK  = "$"
	ARRAY = "*"
)

var (
	PING = "PING"
	GET  = "GET"
	SET  = "SET"
)

var (
	CRLF = "\r\n"
)

type Command struct {
	Cmd  string
	Args []string
	Len  int
}

func NewCommand() *Command {
	return &Command{}
}

func (self *Command) Parse(reader *bufio.Reader) error {
	// read first byte to determine command type
	b, err := reader.ReadByte()
	if err != nil {
		if err == io.EOF {
			return err
		}

		return errors.New("Error reading command : " + err.Error())
	}

	switch string(b) {
	case ARRAY:
		self.Len, err = readLength(reader)
		if err != nil {
			return errors.New("Invalid array length")
		}

		self.Args = make([]string, self.Len-1)
		for i := 0; i < self.Len; i++ {
			b, err := reader.ReadByte()

			// check if the byte is a bulk string
			if err != nil || string(b) != BULK {
				{
					return errors.New("Invalid command format")
				}
			}

			strLen, err := readLength(reader)
			if err != nil {
				return errors.New("Invalid string length")
			}

			// read content
			str := make([]byte, strLen)
			if _, err := io.ReadFull(reader, str); err != nil {
				return errors.New("Error reading string : " + err.Error())
			}

			if _, err := reader.ReadString('\n'); err != nil {
				return errors.New("Expected CRLF")
			}

			if i == 0 {
				self.Cmd = string(str)

				continue
			}

			self.Args[i-1] = string(str)
		}
	default:
		return errors.New("Unknown command type")

	}

	return nil
}

func (self Command) String() string {
	return fmt.Sprintf("Cmd: %s, Args: %v, Len: %d", self.Cmd, self.Args, self.Len)
}

func readLength(reader *bufio.Reader) (int, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	line = line[:len(line)-2]

	length := 0
	for _, ch := range line {
		if ch < '0' || ch > '9' {
			return 0, fmt.Errorf("invalid length")
		}
		length = length*10 + int(ch-'0')
	}
	return length, nil
}
