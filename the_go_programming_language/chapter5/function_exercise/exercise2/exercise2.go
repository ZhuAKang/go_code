package main

import (
	"fmt"
	"time"
	"math/rand"
)

func guessNumber(n int) {
	// 猜数字游戏
	// 传入一个1--100的整数，随机生成一个1--100的整数猜是否和输入的相等
	// 有十次机会，第一次就猜中，提示“你真是个天才” 。。。
	rand.Seed(time.Now().Unix())
	var i int
	for i = 1; i <= 10; i++ {
		if n == rand.Intn(100) + 1 {
			break
		}
	}
	if i == 1 {
		fmt.Println("你真是个天才")
	} else if i <= 3 && i >= 2 {
		fmt.Println("你很聪明，赶上我了")
	} else if i <= 9 && i >= 4 {
		fmt.Println("一般般")
	} else if i == 10{
		fmt.Println("终于猜对了")
	} else {
		fmt.Println("说点啥好呢")
	}

}

func main()  {
	guessNumber(5)
}
