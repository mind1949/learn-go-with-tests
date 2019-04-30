package main

func Sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}

	return sum
}

func SumAll(numbersToSum ...[]int) (sums  []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}

	return sums
}
