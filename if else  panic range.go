package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"os"
	"strings"
)

func main() {
	sop(123165,564,235235,5234234,54234,51,61,)

	//函数作为返回值-----------------------------------
	hello := func() {
		println("hello")

	}
	hello()
	//打印类型
	fmt.Printf("%T", hello)
	//panic---------------------------------------------------
	// 我们使用panic来检查预期不到的错误
	panic("a problem")
	// Panic的通常使用方法就是如果一个函数
	// 返回一个我们不知道怎么处理的错误的
	// 时候，直接终止执行。
	_, err := os.Create("1.txt")
	if err != nil { panic(err) }
	// 键盘输入 --------------------------------------------
	scanner := bufio.NewScanner(os.Stdin)
	//循环获取输入
	for scanner.Scan() {
		//打印 转换为大写
		fmt.Println(scanner.Text() + " || " + strings.ToUpper(scanner.Text())) }
	//检查错误。文件结束不会被当作一个错误
	if err := scanner.Err(); err != nil { fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(0) }
	// range --------------------------------------------
	sopr(1, 2, 4, 6, 8, 10)

	for i, c := range "AB" {
		fmt.Println(i, c)
	}

	//SH1------------------------------------------------
	s1 := sha1.New()

	s1.Write([]byte("ms~go"))

	bs := s1.Sum(nil)

	fmt.Printf("%x", bs)

}
/** 可变参数 */
func sop(args ...int) { for i := 0; i < len(args); i++ { fmt.Println(args[i]) } }
func sopr(args ...int) {
	for _, i := range args {
		fmt.Println(i)
	}}