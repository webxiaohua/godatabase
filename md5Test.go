package main

import (
	"crypto/md5"
	"fmt"
)

func md5String(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func main()  {
	fmt.Println(md5String("123456"))
	var uid int64
	uid = 11110000
	s := fmt.Sprintf("%s_%s","11",string(uid))
	fmt.Println(s)

	data := []byte("123456")
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	fmt.Println(md5str)
}