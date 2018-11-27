package main
import "fmt"
// 变量 字符串
var name = "ms~go"
// 整数
var age = 23
//字符串
var address = "beijing"
// 布尔类型
var sex = true
// 常量
const year = 2018
// 字符
var c = 'c'
func main() {
	fmt.Println(name)
    fmt.Println(age)
	fmt.Println(sex)
    fmt.Println(year)
	var a = 1.5
	var b = 2.5
// 计算加法
   fmt.Println(a + b)
	}
var (
	a int     = 10
	b float64 = 100
)

//a, b := 20, 30

//_, b := 34, 35    丢弃34