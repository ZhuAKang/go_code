package test_03_account

import (
	"fmt"
)

type account struct {
	accountNumber string
	password string
	balance float64
}
// 创建账户
func NewAccount(accountNumber string, password string, balance float64)  *account{
	if len(accountNumber) < 6 || len(accountNumber) > 10 {
		fmt.Println("账号的长度不对...")
		return nil
	}

	if len(password) != 6 {
		fmt.Println("密码的长度不对...")
		return nil
	}

	if balance < 20 {
		fmt.Println("账号的余额不对...")
		return nil
	}

	return &account{
		accountNumber : accountNumber,
		password : password,
		balance : balance,
	}

}

// 存款
func (account *account)Deposite(money float64, password string)  {
	if password != account.password {
		fmt.Println("密码输入错误...")
		return
	}

	if money < 0 {
		fmt.Println("金额输入错误...")
		return
	}

	account.balance += money
	fmt.Println("存款成功...")
}

// 取款
func (account *account)WithDraw(money float64, password string)  {
	if password != account.password {
		fmt.Println("密码输入错误...")
		return
	}

	if money < 0 {
		fmt.Println("金额输入错误...")
		return
	}

	account.balance -= money
	fmt.Println("取款成功...")
}

// 查询余额
func (account *account)Query(password string)  {
	if password != account.password {
		fmt.Println("密码输入错误...")
		return
	}
	fmt.Printf("你的账户 %v 余额为 %v", account.accountNumber, account.balance)
}

// 账户赋值
func (account *account)SetBalance(money float64)  {
	account.balance = money
	return
}
// 账户查询
func (account *account)GetBalance() float64  {
	return account.balance 
}
