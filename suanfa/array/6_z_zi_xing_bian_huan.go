/*
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
*/
package main

func convert(s string, numRows int) string {
	/*
		0   4   8
		1 3 5 7
		2   6

		0     6       12
		1   5 7    11
		2 4   8 10
		3     9

		0       8
		1     7
		2   6
		3 5
		4

		周期t = 2r - 2
		单位周期列数c' = r-1
		最终列数 c= n/t*c' = n/t*(r-1)


	*/
	t := 2*numRows - 1
	c := numRows / t * (numRows - 1)

}
