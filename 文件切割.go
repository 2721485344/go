package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// 获取文件的信息
	fileInfo, _ := os.Stat("27_运算符(上).avi")
	//计算每一个区块的大小
	num := int(math.Ceil(float64(fileInfo.Size()) / float64(chunkSize)))
	//打开文件
	fi, _ := os.OpenFile("27_运算符(上).avi", os.O_RDONLY, os.ModePerm)
	b := make([]byte, chunkSize)
	var i int64 = 1
	for ; i <= int64(num); i++ {
		//移动指针
		fi.Seek((i-1)*(chunkSize), 0)
		if len(b) > int((fileInfo.Size() - (i-1)*chunkSize)) {
			//分配内存
			b = make([]byte, fileInfo.Size()-(i-1)*chunkSize)
			}
		fi.Read(b)
		// 创建块的文件
		f, _ := os.OpenFile(strconv.Itoa(int(i))+".db", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		//写入块文件
		f.Write(b)
		//关闭文件
		f.Close()
		}
	//--------------------------------文件合并--------------------------
	fii, _ := os.OpenFile("1.avi", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	for i := 1; i <= 33; i++ {
		f, err := os.OpenFile(strconv.Itoa(int(i))+".db", os.O_RDONLY, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		b, _ := ioutil.ReadAll(f)
		fii.Write(b)
		f.Close()
		}
    //-----------------------------http 请求----------------
	//生成client 参数为默认
	client := &http.Client{}
	//生成要访问的url
	url := "http://www.baidu.com"
	//提交请求
	reqest, _ := http.NewRequest("GET", url, nil)
	//处理返回结果
	response, _ := client.Do(reqest)
	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	stdout := os.Stdout
	//输出请求体
	_, _ = io.Copy(stdout, response.Body)
	//返回的状态码
	status := response.StatusCode
	//打印输出状态
	fmt.Println(status)

	}


