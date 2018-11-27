package main

import "fmt"

func main()  {
	m:=make(map[string]string)
    m["a"]="a"
	m["b"]="b"
	m["c"]="c"
	//打印字典
	fmt.Println(m)
	//字典长度
	fmt.Println(len(m))
	//根据键取值
	fmt.Println(m["c"])
	//删除
	delete(m,"b")

	fmt.Println(m)
	}
