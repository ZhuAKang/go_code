package main

import (
	"fmt"
)

func main()  {
	// 定义一个打印层数 layer
	var layer int
	fmt.Println("请输入打印的层数：")
	fmt.Scanln(&layer)
	// 定义一个打印的字符--中文字符使用rune类型，
	// 但是由于文字占的大小不一样，打印的空格要改不然格式还是有点问题
	var sty rune
	fmt.Println("请输入打印的字符样式：")
	fmt.Scanf("%c",&sty)


	// 打印基础的金字塔
	for i := 0; i < layer; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%c", sty)
		}
		fmt.Println()
	}

	// 打印高阶金字塔
	for k := 0; k <= layer; k++ {
		
		// 打印空格
		for p := 0; p < layer - k; p++ {
			fmt.Print(" ")
		}

		// 打印字符
		for r := 0; r < 2 * k -1; r++ {
			fmt.Printf("%c", sty)
		}
		fmt.Println()
	}

	// 打印空心金字塔
	for k := 0; k < layer; k++ {
		
		// 打印空格
		for p := 0; p < layer - k; p++ {
			fmt.Print(" ")
		}

		// 打印字符
		for r := 0; r < 2 * k -1; r++ {
			if r == 0 || r == 2 * k - 2  {
				fmt.Printf("%c", sty)
			} else {
				fmt.Print(" ")
			}
			
		}
		fmt.Println()
	}
	for q := 0; q < 2 * layer -1; q++ {
		fmt.Printf("%c", sty)
	}

}

