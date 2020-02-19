/**
Implement a Caesar cipher, both encoding and decoding. The key is an integer from 1 to 25.
This cipher rotates the letters of the alphabet (A to Z).
The encoding replaces each letter with the 1st to 25th next letter in the alphabet (wrapping Z to A).
So key 2 encrypts "HI" to "JK", but key 20 encrypts "HI" to "BC".
This simple "monoalphabetic substitution cipher" provides almost no security, because an attacker who has the encoded message
can either use frequency analysis to guess the key, or just try all 25 keys.
*/
package caesar_cipher

import (
	"fmt"
	"github.com/magmel48/go-projects/utils"
	"strconv"
)

var alphabetLength int64 = 26

func CaesarCipher() {
	text := utils.ReadConsole()

	var rawKey string
	var key int64
	var err error

	for {
		rawKey = utils.ReadConsole()
		key, err = strconv.ParseInt(rawKey, 10, 64)

		if err == nil && key > 0 && key < alphabetLength {
			break
		}
	}

	var out string

	for _, c := range text {
		shifted := c + int32(key)

		if shifted > 'z' {
			shifted = shifted - int32(alphabetLength)
		}

		out = out + string(shifted)
	}

	fmt.Println(out)
}
