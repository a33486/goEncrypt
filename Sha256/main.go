package main

import (
	"crypto/sha256"
	"fmt"
)

func Sha256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs) //将[]byte转成16进制
}

func main() {
	str := "hello world"
	res := Sha256(str)
	fmt.Println(res)
}
