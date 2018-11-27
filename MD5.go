package main

import ( "crypto/md5"
            "fmt"
           "encoding/hex" )
func main() {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte("ms~go"))
	cipherStr := md5Ctx.Sum(nil)
	fmt.Println(cipherStr)
	fmt.Println(hex.EncodeToString(cipherStr))
}