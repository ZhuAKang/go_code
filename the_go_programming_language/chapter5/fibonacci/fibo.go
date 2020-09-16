package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return  1	
	} else{
		return  fibonacci(n - 1) + fibonacci(n - 2)
	}
}

func main() {
	i := 6
	fmt.Printf("第 %d 个Fibonacci数是：%d", i, fibonacci(i))
}