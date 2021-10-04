package main

func Sum(numbers []int) (sum int) {
	sum = 0
	//for i := 0; i < 5; i++ {
	//	sum += numbers[i]
	//}

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numberAll ...[]int) []int {
	//var ret = make([]int, len(numberAll))
	var ret []int
	for _, numbers := range numberAll {
		if len(numbers) <= 0 {
			ret = append(ret, 0)
		} else {
			tail := numbers[1:]
			ret = append(ret, Sum(tail))
		}
	}
	return ret
}

func main() {
	println(Sum([]int{3, 9}))
	var numbers = SumAll([]int{1, 2}, []int{0, 9}, []int{})
	for _, number := range numbers {
		println(number)
	}
}
