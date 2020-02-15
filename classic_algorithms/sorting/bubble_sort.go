package sorting

func BubbleSort(in []int) (out []int) {
	out = in

	for i := 0; i < len(out); i++ {
		found := false

		for j := 0; j < len(out)-i-1; j++ {
			if out[j] > out[j+1] {
				found = true

				t := out[j]
				out[j] = out[j+1]
				out[j+1] = t
			}
		}

		if !found {
			break
		}
	}

	return
}
