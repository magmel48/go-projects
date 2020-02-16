/**
Implement two types of sorting algorithms: Merge sort and bubble sort.
*/
package sorting

func merge(left []int, right []int) (out []int) {
	out = make([]int, len(left)+len(right))

	leftIndex := 0
	rightIndex := 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			out[leftIndex+rightIndex] = left[leftIndex]
			leftIndex++
		} else {
			out[leftIndex+rightIndex] = right[rightIndex]
			rightIndex++
		}
	}

	for i := 0; i < len(left)-leftIndex; {
		out[leftIndex+rightIndex] = left[leftIndex]
		leftIndex++
	}

	for i := 0; i < len(right)-rightIndex; {
		out[leftIndex+rightIndex] = right[rightIndex]
		rightIndex++
	}

	return
}

func sort(in []int) (out []int) {
	if len(in) == 1 {
		return in
	}

	middle := len(in) / 2

	return merge(sort(in[:middle]), sort(in[middle:]))
}

func MergeSort(in []int) (out []int) {
	return sort(in)
}
