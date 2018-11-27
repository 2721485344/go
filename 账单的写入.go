package main
import ( "fmt"
           _ "image/gif"
           _ "image/jpeg"
           _ "image/png"
           "github.com/360EntSecGroup-Skylar/excelize"
            "strconv"
           "io/ioutil"
	"syscall"
)
var i = 1
var xlsx, _ = excelize.OpenFile("1.xlsx")
func main() {
	readDir("F:\\初见文件\\初见\\支付测试")
	xlsx.Save()
	//-----------------文件重命名------------------
	syscall.Rename("新建文本文档.txt","ms.txt")
	//--------------------------环境变量-----------------------
	//获取环境变量
	fs := syscall.Environ()

	//循环输出
	for i = 0; i < len(fs); i++ {
		fmt.Println(fs[i])
	}
	//------------------------------删除文件------------------------
	os.Remove("ms2.txt")
}
func readDir(dirPath string) {
	flist, e := ioutil.ReadDir(dirPath)
	if e != nil {
		fmt.Println("Read file error")
		return
	}
	for _, f := range flist {
		if f.IsDir() {
			readDir(dirPath + "/" + f.Name())
	} else {
			xlsx.AddPicture("Sheet1", "A"+strconv.Itoa(i), dirPath+"/"+f.Name(), `{"x_scale": 0.5, "y_scale": 0.5}`)
			i += 49
			}
	}
	}
