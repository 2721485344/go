package main

import "fmt"

func main() {
	var a = 0xff
	var b = 10
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)
	fmt.Println(a >> 1)
	fmt.Println(b << 1)
	fmt.Println(b | 1)
	fmt.Println(b & 1)
	fmt.Println(b != 1)

}
