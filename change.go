package main

import (
	"fmt"
	"math"
)

func change(values []int, total int) []int {
	result := make([]int, total+1)
	temp := make([]int, total+1)
	temp[0] = 0

	for i := 1; i <= total; i++ {
		temp[i] = math.MaxInt32

		for j := 0; j < len(values); j++ {
			if i >= values[j] && temp[i-values[j]]+1 < temp[i] {
				temp[i] = temp[i-values[j]] + 1
				result[i] = j + 1
			}
		}
	}

	return result
}

func printChange(values []int, total int) {
	S := change(values, total)

	for total > 0 {
		fmt.Println(values[S[total]-1])
		total = total - values[S[total]-1]
	}
}

func main() {
	examples := []int{
		15,
		48,
		50,
		64,
		74,
		98,
		108,
		126,
		220,
		406,
		486,
	}
	total := 2445

	fmt.Printf("Examples:\n%v\n\n", examples)
	fmt.Printf("Total %d\n\n", total)
	fmt.Println("Result")
	printChange(examples, total)

	fmt.Scanln()
}
