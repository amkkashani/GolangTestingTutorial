package main

import (
	"fmt"
	"time"
)

func main() {
	test()
	time.Sleep(5 * time.Second)
}

func Summer(x,y int) int  {
	if x <0 {
		x = -x
	}

	if y < 0{
		y = -y
	}
	return x + y
}

var tests = []int{100,55,99,45,21,654,87,21,55,88,99,11,1,2,3,45,5,6,78,9,10}


func test()  {
	for _, tc := range tests {
		temp := tc // capture range variable
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Println(temp)
		}()
	}
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

