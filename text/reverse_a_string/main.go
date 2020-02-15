/**
Enter a string and the program will reverse it and print it out.
*/
package reverse_a_string

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(in string) (out string) {
	for _, c := range in {
		out = string(c) + out
	}

	return
}

func ReverseAString() {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	text = text[:len(text)-1]

	fmt.Println(reverse(text))
}
