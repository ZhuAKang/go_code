package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	l := len(s)
	if l <= 3 {
		return s
	} else {
		buf.WriteByte(s[0])
		for i :=1 ; i < l; i++ {
			if (l - i) % 3 == 0 {
				buf.WriteByte(',')
				buf.WriteByte(s[i])
			} else {
				buf.WriteByte(s[i])
			}
		}
		return buf.String()

	}
}

func main() {
	fmt.Println(comma("123456"))
	a := 1
	b := 2
	a ,b = b, a
	fmt.Printf("a = %d; b = %d \n", a, b)
	c  := -1>>2
	fmt.Printf("%d",c)
}