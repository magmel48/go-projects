/**
Enter a string and the program will reverse it and print it out.
*/
package reverse_a_string

import (
	"fmt"
	"github.com/magmel48/go-projects/utils"
)

func Reverse(in string) (out string) {
	for _, c := range in {
		out = string(c) + out
	}

	return
}

func ReverseAString() {
	text := utils.ReadConsole()
	fmt.Println(Reverse(text))
}
