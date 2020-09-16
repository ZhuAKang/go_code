package main

import (
	"fmt"
	"os"
	"time"
	"strings"
)

func main()  {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "		
	}
	fmt.Println(s)
	fmt.Println("循环叠加耗时")
	fmt.Printf("%.7fs \n", time.Since(start).Seconds())
	start = time.Now()
	s = strings.Join(os.Args[1:]," ")
	fmt.Println(s)
	fmt.Println("使用join耗时")
	fmt.Printf("%.7fs \n", time.Since(start).Seconds())


}