package main

import (
	"fmt"
	"os"
)
func main()  {
	t := os.Args[1]
	n := int(t)
	fmt.Printf("t 的数据格式为：%T", n)
}