package main

import (
	"fmt"
	"math"
)

func suShu(n int) {
	var flag bool
	fmt.Printf("%d ", 2)
	for i := 3; i < n; i +=2 {
		flag = true
		for j := 3; float64(j) < math.Sqrt(float64(i)); j++ {	
			if i % j == 0 {
				flag = false
				break
			}		
		}
		if flag {
				fmt.Printf("%d ", i)
			}		
	} 
}

func main() {
	suShu(100)

}