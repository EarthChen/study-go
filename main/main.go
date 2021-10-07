package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

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

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	Greet(os.Stdout, "test")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
	//println(Sum([]int{3, 9}))
	//var numbers = SumAll([]int{1, 2}, []int{0, 9}, []int{})
	//for _, number := range numbers {
	//	println(number)
	//}
}
