package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

func main(){
	str := "[1,2,3,4,5,6,7,8,9,0,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30]" // 81 74
	//str := "[1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]"  // 41 30
	strByte := []byte(str)
	fmt.Println("len1:",len(strByte))
	var buf bytes.Buffer
	err:= gzipConfig(&buf,strByte)
	if err!=nil {
		fmt.Println("error:",err)
	}else{
		fmt.Println("len2:",len(buf.Bytes()))
	}
}

func gzipConfig(buf *bytes.Buffer,conf []byte)error{
	w := gzip.NewWriter(buf)
	defer w.Close()
	if _,err := w.Write(conf);err != nil {
		return err
	}
	return nil
}