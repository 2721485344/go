package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

var start = 0
var len = start + 10
func main() {
	for {
		var url = "http://pic.sogou.com/pics/channel/getAllRecomPicByTag.jsp?category=%E5%A3%81%E7%BA%B8&tag=%E5%85%A8%E9%83%A8&start=" + strconv.Itoa(start) + "&len=" + strconv.Itoa(len)
		fmt.Println(url)
		res, _ := http.Get(url)
		reader, _ := ioutil.ReadAll(res.Body)
		var filename = strconv.Itoa(start) + ".json"
		os.Create(filename)
		ioutil.WriteFile(filename, reader, 0777)
		start = len
		len = start + 10
		if len > 1024 {
			break
		}
		}
}
