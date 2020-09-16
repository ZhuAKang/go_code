package main

import (
	"fmt"
	"runtime"
)

func main()  {
	cpuNum := runtime.NumCPU()
	fmt.Printf("CPU的个数为%d \n", cpuNum)

	// 设置自己go程序使用多少CPU
	runtime.GOMAXPROCS(cpuNum)

	//
	data := make(chan int, 3)
	data<- 1
	data<- 2
	data<- 3
	data1, ok := <-data
	fmt.Printf("是否可以取出: %v ; 取出的数据是：%d \n", ok, data1)
	close(data)
	data2, ok := <-data
	fmt.Printf("关闭后是否可以取出: %v ; 取出的数据是：%d \n", ok, data2)
	<-data //取光数据
	data4, ok := <-data
	fmt.Printf("channel空了后是否可以取出: %v ; 取出的数据是：%d \n", ok, data4)
}