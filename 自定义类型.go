package main

import "fmt"

// 自定义类型
type ms string
//结构体
type perwson struct {
	name ms
}
func main() {
	var p1 = new(perwson)
	p1.name="ms"
	fmt.Println(p1)
	}
