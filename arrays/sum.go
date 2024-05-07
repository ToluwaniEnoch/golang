package main

func Sum(numbers [5]int) int {
	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}

func SumOfSlice(numbers []int) int {
	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}

func SumAll (numbersToSum ...[]int) []int {
	return returnSum(numbersToSum...)

}

func SumAllTails (numbersToSum ...[]int) []int {
	tail := numbersToSum[1:]
	return returnSum(tail...)
}

func returnSum(numbersToSum ... []int) []int {
	var result []int
	
	if len(numbersToSum) == 0 {
		result = append(result, 0)
	
	} else {
		for _, num := range numbersToSum {
			result = append(result,  SumOfSlice(num))
		}
	}
	
	return result
}