package main

import (
	"github.com/magmel48/go-projects/text/count_words_in_a_string"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	//closest_pair_problem.ClosestPairProblem()
	count_words_in_a_string.CountWordsInAString()
}
