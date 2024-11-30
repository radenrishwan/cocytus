package cocytus

import "bufio"

func WriteCommand(writer bufio.Writer, cmd string, args []string) error {
	r := ""

	r += cmd
	if len(args) > 0 {
		for _, arg := range args {
			r += " " + arg
		}
	}

	_, err := writer.WriteString(r)
	if err != nil {
		return err
	}

	return writer.Flush()
}
