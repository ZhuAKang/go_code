package main

import (
	"fmt"
)

func printDays() {
	// 循环打印输入的月份的天数（使用 continue 实现） 	
	// 要有判断输入月份是否错误的语句
	fmt.Println("欢迎，输入-1退出系统")
	defer fmt.Println("系统退出成功！")
	var year int
	var month int
	for  {
		fmt.Println("请输入年份：")
		fmt.Scanln(&year)
		fmt.Println("请输入月份：")
		fmt.Scanln(&month)
		if month == -1 || year == -1 {
			break 
		}
		switch month {
			case 1, 3, 5, 7, 8, 10, 12:
				fmt.Printf("%d 月有 31 天 \n", month)
			case 4, 6, 9, 11:
				fmt.Printf("%d 月有 30 天 \n", month)
			case 2:
				if year % 4 == 0 {
					fmt.Printf("%d 年 %d 月有 28 天 \n", year, month)
				} else {
					fmt.Printf("%d 年 %d 月有 29 天 \n", year, month)
				}	
			default:
				fmt.Println("输入有误...")
		}

	}


}

func main()  {
	printDays()

}