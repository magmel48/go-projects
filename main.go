package main

import (
	"fmt"
	"github.com/magmel48/go-projects/classic_algorithms/sorting"
)

func main() {
	//fizz_buzz.FizzBuzz()
	//reverse_a_string.ReverseAString()
	//pig_latin.PigLatin()
	//count_vowels.CountVowels()
	fmt.Println(sorting.BubbleSort([]int{-1, 0, -10, 9, 8, 4, 2, -3, 8}))
	fmt.Println(sorting.MergeSort([]int{-1, 0, -10, 9, 8, 4, 2, -3, 8}))
}
