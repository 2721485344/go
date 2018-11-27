package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// 这个"intSeq"函数返回另外一个在intSeq内部定义的匿名函数，
// 这个返回的匿名函数包住了变量i，从而形成了一个闭包
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
		}
	}
func main() {
	// 我们调用intSeq函数，并且把结果赋值给一个函数nextInt，
	// 这个nextInt函数拥有自己的i变量，这个变量每次调用都被更新。
	// 这里i的初始值是由intSeq调用的时候决定的。
	nextInt := intSeq()
	// 调用几次nextInt，看看闭包的效果
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	// 为了确认闭包的状态是独立于intSeq函数的，再创建一个。
	newInts := intSeq()
	fmt.Println(newInts())
	//-------------------------o创建文件夹----------
	os.Mkdir("hello", 0777)

    //----------------------创建1024个文件夹,然后在删除-----------------
	for i := 0; i < 1024; i++ {
		os.Mkdir("hello"+strconv.Itoa(i), 0777)
	}
	time.Sleep(time.Second * 30)
	for i := 0; i < 1024; i++ {
		os.Remove("hello" + strconv.Itoa(i))
	}

	}
