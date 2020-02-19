package closest_pair_problem

import (
	"fmt"
	"testing"
)

func TestPlanarCase(t *testing.T) {
	dataset := []point{
		{1, 2},
		{3, 4},
		{5, 6},
		{1, 0},
		{1, 1},
	}

	d, pair := planarCase(dataset)

	if d != 1 || pair[0].x != 1 {
		fmt.Println(d, pair)
		t.Fail()
	}
}
