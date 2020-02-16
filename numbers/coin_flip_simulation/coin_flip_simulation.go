/**
Write some code that simulates flipping a single coin however many times the user decides.
The code should record the outcomes and count the number of tails and heads.
*/
package coin_flip_simulation

import (
	"fmt"
	"github.com/magmel48/go-projects/utils"
	"math/rand"
	"strconv"
)

func CoinFlipSimulation() {
	var rawTimes string
	var times int64
	var err error

	for {
		rawTimes = utils.ReadConsole()
		times, err = strconv.ParseInt(rawTimes, 10, 64)

		if err == nil && times != 0 {
			break
		}
	}

	tails := 0
	heads := 0

	for i := 0; i < int(times); i++ {
		n := rand.Float64()

		if n < 0.5 {
			tails++
			fmt.Println("tail")
		} else {
			heads++
			fmt.Println("head")
		}
	}

	fmt.Println("\n", "\bsummary:")
	fmt.Println("tails:", tails)
	fmt.Println("heads:", heads)
}
