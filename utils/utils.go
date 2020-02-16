package utils

import (
	"bufio"
	"os"
)

func ReadConsole() (out string) {
	reader := bufio.NewReader(os.Stdin)
	out, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	out = out[:len(out)-1]

	return
}
