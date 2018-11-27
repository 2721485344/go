package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 创建文件
	f, _ := os.Create("hello world.txt")
	//defer关闭文件
	defer f.Close()
	//写入文件 方法1
	f.Write([]byte("hello world "))
	//写入文件 方法2
	f.WriteString("你是我的眼")
	//同步文件
	f.Sync()
	//写入文件方法3
	w := bufio.NewWriter(f)
	//写入数据
	w.WriteString("哈哈哈哈哈哈")
	//刷新
	w.Flush()
	// 文件写入-------------------------------------
	var str = "File Writer"
	ioutil.WriteFile("fw.txt", []byte(str), 0777)
	// 文件读取----------------------------------------
	fr, _ := ioutil.ReadFile("fw.txt")
	fmt.Println(string(fr))

	}
