package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func read0(path string) string {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%s\n", err)
		panic(err)
	}
	return string(f)
	}
func read1(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for { n, err := fi.Read(buf)
	if err != nil && err != io.EOF {
		panic(err)
	}
	if 0 == n {
		break
	}
	chunks = append(chunks, buf[:n]...)
	}
	return string(chunks)
	}
func read2(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)
	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
	if err != nil && err != io.EOF {
		panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
		}
	return string(chunks) }
func read3(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
	}
func main() {
	str, _ := http.Get("http://pic.sogou.com/pics/channel/getAllRecomPicByTag.jsp?category=%E5%A3%81%E7%BA%B8&tag=%E5%85%A8%E9%83%A8&start=0&len=100000")
	html, _ := ioutil.ReadAll(str.Body)
	ioutil.WriteFile("fw.json", html, 0777)
	file := "fw.json"
	start := time.Now()
	read0(file)
	t0 := time.Now()
	fmt.Printf("Cost time %v\n", t0.Sub(start))
	read1(file)
	t1 := time.Now()
	fmt.Printf("Cost time %v\n", t1.Sub(t0))
	read2(file)
	t2 := time.Now()
	fmt.Printf("Cost time %v\n", t2.Sub(t1))
	read3(file)
	t3 := time.Now()
	fmt.Printf("Cost time %v\n", t3.Sub(t2))
	//-----------------------递归遍历目录-----------------------
	//GO语言获取目录列表用 ioutil.ReadDir()，遍历目录用 filepath.Walk()
	}
