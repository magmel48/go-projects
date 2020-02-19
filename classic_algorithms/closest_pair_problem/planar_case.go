/**
inspired by https://www.geeksforgeeks.org/closest-pair-of-points-using-divide-and-conquer-algorithm/
*/
package closest_pair_problem

import (
	"math"
	"sort"
)

func closestInStrip(data []point, min float64) (length float64, closestPair []point) {
	length = min

	out := make([]point, len(data))
	copy(out, data)

	sort.Slice(out, func(i, j int) bool {
		return out[i].y < out[j].y
	})

	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out) && (data[j].y-data[i].y) < length; j++ {
			if data[i].distance(data[j]) < length {
				length = data[i].distance(data[j])
				closestPair = []point{data[i], data[j]}
			}
		}
	}

	return
}

func divide(data []point) (length float64, closestPair []point) {
	if len(data) <= 3 {
		return bruteForce(data)
	}

	middle := len(data) / 2
	middlePoint := data[middle]

	dl, inLeft := divide(data[:middle])
	dr, inRight := divide(data[middle:])

	d := math.Min(dl, dr)

	strip := make([]point, 0)
	for _, el := range data {
		if math.Abs(el.x-middlePoint.x) < d {
			strip = append(strip, el)
		}
	}

	ds, inStrip := closestInStrip(strip, d)

	length = ds
	closestPair = inStrip

	if dl < dr && dl < ds {
		length = dl
		closestPair = inLeft
	} else if dr < ds {
		length = dr
		closestPair = inRight
	}

	return
}

func planarCase(data []point) (length float64, closestPair []point) {
	out := make([]point, len(data))
	copy(out, data)

	sort.Slice(out, func(i, j int) bool {
		return out[i].x < out[j].x
	})

	return divide(out)
}
