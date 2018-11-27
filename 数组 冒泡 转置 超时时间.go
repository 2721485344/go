package main


import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 常量数组长度
	const N = 10
	// 声明数组
	var arr = [N]int{}
	// 随机数
	 r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 循环赋值
	 for i := 0; i < len(arr); i++ {
	// 赋值
	arr[i] = r.Intn(0xff)
	}
	// 打印数组
	fmt.Println(arr)
	 /*******************************************
	                      冒泡排序
	 ********************************************/
	// :=自动匹配变量类型
	 num := len(arr)
	//花括号{必须在这一行 if也一样
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ { if arr[i] > arr[j] {
			//相比某语言不需要临时交换变量
			arr[i], arr[j] = arr[j], arr[i]
			}
		}
		}
	//输出排序后的数组
	fmt.Println(arr)
   //----------------------------------------数组和冒泡-----------
	arr1 := []int{10, 12, 33, 14, 65}
	for _, i := range arr {
		fmt.Printf("%d ",i)
	}
	fmt.Println()
	bubblesort(arr1)
	for _, i := range arr {
		fmt.Printf("%d ",i)
	}
   //------------------------一维数组,二位数组,二维数组的转置---------------------------------------
	var a [5] int
	a[0] = 1024 >> 10
	fmt.Printf("%d \n", a)
	var c = 0
	var b [3][3]int
	var d [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c++
			b[i][j] = c
			}
		}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", b[i][j])
		}
		fmt.Println()
		}
	fmt.Println("----------")
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[0]); j++ {
			d[i][j] = b[j][i]
			}
		}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", d[i][j])
		}
		fmt.Println()
		}
	//----------------------------超时时间---------------------
	c1 := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c1:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
			o <- true
			break
			}
	}
	}()
	<-o

}

func bubblesort(n[]int) {
	for i := 0; i < len(n)-1; i++ {
	for j := i + 1; j < len(n); j++ {
		if n[j] < n[i] {
			n[i], n[j] = n[j], n[i]
			}
	}
	}
	}
