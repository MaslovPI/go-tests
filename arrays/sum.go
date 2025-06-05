package main

func Sum(array []int) int {
	sum := 0
	for _, number := range array {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, slice := range numbersToSum {
		sums = append(sums, Sum(slice))
	}
	return sums
}

func SumAllTail(numbersToSum ...[]int) []int {
	var sums []int
	for _, slice := range numbersToSum {
		if len(slice) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail := slice[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
