package main

import (
	"encoding/base64"
	"fmt"
)

//base64加密
func Base64Encode(str string) string {
	strbytes := []byte(str)
	return base64.StdEncoding.EncodeToString(strbytes)
}

//base64解密
func Base64Decode(str string) string {
	decoded, _ := base64.StdEncoding.DecodeString(str)
	return string(decoded)

}

func main() {
	var str = "hello world"
	encoded := Base64Encode(str)
	fmt.Println(encoded)
	decodestr := Base64Decode(encoded)
	fmt.Println(decodestr)
}
