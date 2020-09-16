package main

import (
	"fmt"
)

type Student struct {
	name string
	gender string
	age int
	id int
	score float64
}

func (student *Student)say() string {
	str := fmt.Sprintf("student 的信息 name = [%v] , gender = [%v], age = [%v], id = [%v], score = [%v]",
	 student.name, student.gender, student.age, student.id, student.score)
	return str 
}

func main() {
	stu := Student{
		name : "tom",
		gender : "男",
		age : 18,
		id : 1000,
		score : 66.66,
	}

	fmt.Println(stu.say())
}
