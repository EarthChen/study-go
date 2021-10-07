package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
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

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, "Go!")
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

func main() {
	Countdown(os.Stdout, &ConfigurableSleeper{1 * time.Second})
	//Greet(os.Stdout, "test")
	//http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
	//println(Sum([]int{3, 9}))
	//var numbers = SumAll([]int{1, 2}, []int{0, 9}, []int{})
	//for _, number := range numbers {
	//	println(number)
	//}
}
