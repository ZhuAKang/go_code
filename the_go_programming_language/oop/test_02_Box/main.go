package main

import (
	"fmt"
)

type Box struct {
	len float64
	height float64
	width float64
}

func (box *Box)getVolume() float64 {
	return box.len * box.height * box.width
}

func main() {
	var box Box
	fmt.Print("请输入箱子长度：")
	fmt.Scanln(&(box.len))
	fmt.Print("请输入箱子高度：")
	fmt.Scanln(&(box.height))
	fmt.Print("请输入箱子宽度：")
	fmt.Scanln(&(box.width))

	fmt.Printf("箱子的体积为：%f", box.getVolume())
}