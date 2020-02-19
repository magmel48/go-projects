/**
Checks if the string entered by the user is a palindrome. That is that it reads the same forwards as backwards like “racecar”
*/
package check_if_palindrome

import (
	"fmt"
	"github.com/magmel48/go-projects/text/reverse_a_string"
	"github.com/magmel48/go-projects/utils"
	"strings"
)

func CheckIfPalindrome() {
	text := strings.Replace(utils.ReadConsole(), " ", "", -1)
	reversed := reverse_a_string.Reverse(text)

	if text == reversed {
		fmt.Println("it's palindrome")
	} else {
		fmt.Println("it's just ordinary phrase")
	}
}
