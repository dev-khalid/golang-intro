package functionsdeepdive

import "fmt"

func sum(numbers ...int) int {
	sum := 0
	for value := range numbers {
		sum += value
	}

	return sum
}

func VariadicFunction() {
	nums := []int{1,2,3,4,5}

	fmt.Println(sum(nums...))
}