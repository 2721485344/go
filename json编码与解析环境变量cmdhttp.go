package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

type Person struct {
	Name string
	Age int
	Sex bool
}
func main() {
	/*****json转换*****/
	//创建对象
	person := Person{
		"ms~go",
		23,
		true,
	}
	//json编码
	personjson, _ := json.Marshal(person)
	//输出json字符串
	fmt.Println(string(personjson))
	/*****json解析*****/
	strjson := `{"Name":"ms~go","Age":23,"Sex":true}`
	r := &Person{}
	json.Unmarshal([]byte(strjson), &r)
	fmt.Println(r)
	fmt.Println(r.Name)
	fmt.Println(r.Age)
	fmt.Println(r.Sex)
	//查看某个环境变量-----------------------------------
	ANDROID_HOME,AB:=os.LookupEnv("ANDROID_HOME")
	fmt.Println(ANDROID_HOME)
	fmt.Println(AB)
	JAVA_HOME,JB:=os.LookupEnv("JAVA_HOME")
	fmt.Println(JAVA_HOME)
	fmt.Println(JB)
	PATH,PB:=os.LookupEnv("PATH")
	fmt.Println(PATH)
	fmt.Println(PB)
	//执行CMD命令-------------------------------------------------
	cmdcalc := exec.Command("calc")
	cmdcalc.Start()
	cmdnotepad := exec.Command("notepad")
	cmdnotepad.Start()
	cmdmspaitn := exec.Command("mspaint")
	cmdmspaitn.Start()
	//http 请求-------------------------------------------------------
	resp, _ := http.Get("http://www.baidu.com/")
	//defer 打开
	defer resp.Body.Close()
	//获取请求体
	body, _ := ioutil.ReadAll(resp.Body)
	//打印请求到的数据
	fmt.Println(string(body))
	//创建文件
	os.Create("baidu.html")
	//写入数据
	ioutil.WriteFile("baidu.html", body, 777)

	}
