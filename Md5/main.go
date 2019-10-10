package main

import (
	"crypto/md5"
	"fmt"
)

func toMd5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}

func main() {
	str := "hello world"
	res := toMd5(str)
	fmt.Println(res)
}
