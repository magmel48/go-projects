/**
Pig Latin is a game of alterations played on the English language game.
To create the Pig Latin form of an English word the initial consonant sound is transposed to the end of the word
and an ay is affixed (Ex.: "banana" would yield anana-bay). Read Wikipedia for more information on rules.
*/
package pig_latin

import (
	"bufio"
	"fmt"
	"os"
)

var vowels = []string{"a", "o", "e", "i", "u"}
var pigAddition = "ay"

func IsVowel(in string) (out bool) {
	for _, v := range vowels {
		if string(in) == v {
			return true
		}
	}

	return
}

func PigLatin() {
	reader := bufio.NewReader(os.Stdin)

	var text string
	var err error

	for len(text) == 0 {
		text, err = reader.ReadString('\n')

		if err != nil {
			panic(err)
		}

		text = text[:len(text)-1]
	}

	formatted := false

	if IsVowel(string(text[0])) {
		text = text + "-" + pigAddition
		formatted = true
	}

	if !formatted {
		consonantsCount := 0

		for _, c := range text {
			if IsVowel(string(c)) {
				break
			}

			consonantsCount++
		}

		text = text[consonantsCount:] + "-" + text[:consonantsCount] + pigAddition
	}

	fmt.Println(text)
}
