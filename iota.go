package main
import "fmt"
func main() {
	//为 0
	const a = iota
	const (
	// 没增加一行+1
	b = iota
	c = iota
	d = iota
	e, f = iota, iota
	// 同一行数值相等
	g, h, i, j, k = iota, iota, iota, iota, iota )
	fmt.Printf("a = %d , b = %d , c = %d , d = %d , e = %d , f = %d , g = %d , h = %d , i = %d , j = %d , k = %d ", a, b, c, d, e, f, g, h, i, j, k)
	//每次 const 出现时，都会让 iota 初始化为0.
	const l = iota
	fmt.Println(l)

	// 自增长类型
	//--------------------------------------------
	//const (
	//	a = iota
	//	b
	//	c
	//	d
	//)
	//--------------------------------------------
	//使用 “_” 可以跳过值
	//	const (
	//		a = iota
	//		b
	//		c
	//		// 使用 "_" 可以跳过值
	//		_
	//		_
	//		d
	//	)
	//--------------------------------------------
	//位掩码表达式
	//const(
	//	a=1<<iota b
	//c
	//d
	//e
	//f
	//_
	//g
	//h )
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	//fmt.Println(d)
	//fmt.Println(e)
	//fmt.Println(f)
	//fmt.Println(g)
	//fmt.Println(h)
	//定义数量级别
	//type ByteSize float64
	//const ( _ = iota  // ignore first value by assigning to blank identifier
	//	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	//	MB // 1 << (10*2)
	//	GB // 1 << (10*3)
	//	TB // 1 << (10*4)
	//	PB // 1 << (10*5)
	//	EB // 1 << (10*6)
	//	ZB // 1 << (10*7)
	//	YB // 1 << (10*8) )
	}


