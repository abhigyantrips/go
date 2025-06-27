package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(vectors ...[]int) []int {
	var sums []int
	
	for _, numbers := range vectors {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(vectors ...[]int) []int {
	var sums []int

	for _, numbers := range vectors {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}