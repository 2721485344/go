package main
import "fmt"
func main() {
	// 声明变量
	var a int = 10
	var b int = 10
	// 打印地址
	fmt.Println("pointer address : %x", &a)
	fmt.Println("pointer address : %x", &b)
	// 判断地址是否相同
	fmt.Println(&a == &b)
	// 值赋值
	var c = a
	// 值赋值
	var d = b
	// 打印值
	fmt.Println(c)
	// 打印地址
	fmt.Println("pointer address : %x", &c)
	// 打印值
	fmt.Println(d)
	// 打印地址
	fmt.Println("pointer address : %x", &d)
	// 取地址
	var e = &a
	// 取地址
	var f = &a
	// 输出地址
	fmt.Println("pointer address : %x", &e)
	// 输出地址
	fmt.Println("pointer address : %x", &f)
	// 输出值
	fmt.Println(e)
	// 输出值
	fmt.Println(f)
	// 判断地址是否相同
	fmt.Println(e == f)
	// 指针类型
	var ip *int = &a
	// 打印指针
	fmt.Println(ip)
	// 二级指针
	var jp **int = &ip
	fmt.Println(jp)
	// 三级指针
	var kp ***int = &jp;
	fmt.Println(kp)
	arr := []int{10, 100, 200}
	var i int
	for i = 0; i < len(arr); i++ {
		fmt.Printf("a[%d] = %d | %x \n", i, arr[i], &arr[i])
	}
	}
