/**
The closest pair of points problem or closest pair problem is a problem of computational geometry:
given n points in metric space, find a pair of points with the smallest distance between them.
*/
package closest_pair_problem

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

func (p1 point) distance(p2 point) (out float64) {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

func bruteForce(data []point) (length float64, closestPair []point) {
	length = math.Inf(1)

	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] != data[j] {
				d := data[i].distance(data[j])

				if d < length {
					length = d
					closestPair = []point{data[i], data[j]}
				}
			}
		}
	}

	return
}

func ClosestPairProblem() {
	points := []point{
		{1, 2},
		{3, 4},
		{0, 9},
		{1.1, 2},
		{1.2, 2},
		{2, 2},
	}

	fmt.Println(bruteForce(points))
	fmt.Println(planarCase(points))
}
