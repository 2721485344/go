package main

import (
	"fmt"
	"time"
)

func main() {
	//必须先运行ms和ms1
	// 如果先运行for{}
	// 则执行不到 ms() ms1()
	go ms0()
	go ms1()
	for {
		fmt.Println("this is main ")
		time.Sleep(time.Second) }
}
// 函数1
func ms0() {
	for {
		fmt.Println("this is ms")
		time.Sleep(time.Second)
}
}
//函数2
func ms1() {
	for {
		fmt.Println("this is ms1")
		time.Sleep(time.Second)
}
}
