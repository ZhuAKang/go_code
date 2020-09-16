package main

import (

	"fmt"
	"go_code/the_go_programming_language/oop/test_03_account"
)

func main()  {

	account := test_03_account.NewAccount("adsadadad", "565656", 52)

	if account != nil {

		fmt.Println("创建成功...")

	} else {

		fmt.Println("创建失败...")

	}

	account.SetBalance(100)
	fmt.Printf("当前账户的余额是 %f ", (*account).GetBalance())
}