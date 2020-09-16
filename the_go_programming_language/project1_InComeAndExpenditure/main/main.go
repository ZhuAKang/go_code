package main

import (
	"fmt"
	"strconv"
)

func checkIncomeAndExpenditure(ledger []map[int]map[string]string)  {
	fmt.Println("------------当前收支明细----------------")
	for i := 0; i < len(ledger); i++ {
		
		fmt.Println(ledger[i][i]["收支"], ledger[i][i]["账户金额"], ledger[i][i]["收支金额"], ledger[i][i]["说明"])
	}
}

func Income(ledger []map[int]map[string]string)  {
	var newIncome map[int]map[string]string
	newIncome = make(map[int]map[string]string, 1)
	newIncome[len(ledger)] = make(map[string]string)
	newIncome[len(ledger)]["收支"] = "收入" 
	// 获取收支
	var incomeMoney int
	var details string
	fmt.Print("输入收入的金额：")
	fmt.Scan(&incomeMoney)
	fmt.Print("输入收入的说明：")
	fmt.Scan(&details)

	// 写入收支
	strMoney := ledger[len(ledger) - 1][len(ledger) - 1]["账户金额"]
	money , _ := strconv.Atoi(strMoney)
	money += incomeMoney
	newIncome[len(ledger)]["账户金额"] = strconv.Itoa(money) 
	newIncome[len(ledger)]["收支金额"] = strconv.Itoa(incomeMoney)
	newIncome[len(ledger)]["说明"] = details
	ledger = append(ledger, newIncome)
}

func main()  {
	// 用来保存账本的 map 切片
	var ledger []map[int]map[string]string
	ledger = make([]map[int]map[string]string, 1)
	// 第一笔存在ledger[0]中，包含map[0]{
	//                                  map{收支：收入/支出
	//	                                    账户金额：1000
	//                                      收支金额：500
	//                                      说明：买菜
	//                                     }
	//                                 }
	// 初始化收支
	ledger[0] = make(map[int]map[string]string)

	for {
		fmt.Println("-----------------------------------------------")
		fmt.Println("------------------家庭收支记账软----------------")
		fmt.Println("                  1 收支明细                    ")
		fmt.Println("                  2 登记收入                    ")
		fmt.Println("                  3 登记支出                    ")
		fmt.Println("                  4 退    出                    ")
		fmt.Print("                  请选择<1-4>：")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("收支明细")
			checkIncomeAndExpenditure(ledger)
		case 2:
			fmt.Println("登记收入")
			Income(ledger)
		case 3:
			fmt.Println("登记支出")
		case 4:
			fmt.Println("退出")
			return
		default:
			fmt.Println("请输入正确的选项")
		}


	}
}