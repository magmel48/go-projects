/**
Ask the user to enter 2 integers a and b and output a^b (i.e. pow(a,b)) in O(lg n) time complexity.
*/
package fast_exponentiation

import (
	"fmt"
	"github.com/magmel48/go-projects/utils"
	"strconv"
	"strings"
)

func squaring(in float64, n int64) (out float64) {
	if n < 0 {
		return squaring(1/in, -n)
	} else if n == 0 {
		return 1
	} else if n == 1 {
		return in
	} else if n%2 == 0 {
		return squaring(in*in, n/2)
	}

	return in * squaring(in*in, (n-1)/2)
}

func FastExponentiation() {
	var rawInput []string
	var a int64
	var b int64
	var err error

	for {
		rawInput = strings.Split(utils.ReadConsole(), " ")

		if len(rawInput) == 2 {
			a, err = strconv.ParseInt(rawInput[0], 10, 64)
			if err == nil && a != 0 {
				b, err = strconv.ParseInt(rawInput[1], 10, 64)
				if err == nil && b != 0 {
					break
				}
			}
		}
	}

	fmt.Println(squaring(float64(a), b))
}
