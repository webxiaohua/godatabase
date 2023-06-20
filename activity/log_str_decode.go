package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// 转换8进制utf-8字符串到中文
// eg: `\346\200\241` -> 怡
func convertOctonaryUtf8(in string) string {
	s := []byte(in)
	reg := regexp.MustCompile(`\\[0-7]{3}`)

	out := reg.ReplaceAllFunc(s,
		func(b []byte) []byte {
			i, _ := strconv.ParseInt(string(b[1:]), 8, 0)
			return []byte{byte(i)}
		})
	return string(out)
}

func main() {
	s1 := "\347\263\273\347\273\237\347\271\201\345\277\231\357\274\214\350\257\267\347\250\215\345\220\216\345\206\215\350\257\225" // 字面量

	// 转化 s2
	s3 := convertOctonaryUtf8(s1)
	fmt.Println(s3)
}
