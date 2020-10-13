package video_code

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func stringAndFile() {
	fi, err := os.Open("D:\\go\\goProject\\src\\go_code\\code\\video_code\\数据.txt")
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	defer fi.Close() // 延迟关闭文件
	br := bufio.NewReader(fi)
	for {
		line, _, err := br.ReadLine()
		// readline 返回的 line 是一个字符数组
		// readline 较为初级，大多数调用者应使用ReadBytes('\n')或ReadString('\n')代替
		// line, err := br.ReadString('\n')
		// line, err := br.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		str := string(line)
		// 字符串切分
		strs := strings.Split(str, "#")
		fmt.Println(strs[1])
	}
	// TODO: 上面这个操作涉及到磁盘，所以在执行搜索排序等功能的时候磁盘操作较慢，可以先将数据放入到内存中
}
