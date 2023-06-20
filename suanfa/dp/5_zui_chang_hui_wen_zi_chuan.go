// 5. 最长回文子串
/*
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
*/
package main

import "fmt"

func main() {
	res := longestPalindrome("aaaa")
	fmt.Println(res)
}
func longestPalindrome(s string) string {
	// 判断是否为回文串  d[i][j]  从i~j是否为回文串
	// i==j  true
	// i与j 相差1， 例如：aa 也是回文串
	// s[i] == s[j] true
	// s[i] != s[j] false
	// i++ j-- d[i+1][j-1]
	// dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
	cnt := len(s)
	if cnt < 2 {
		return s
	}
	maxLen := 1
	begin := 0
	dp := make([][]bool, cnt)
	for i := 0; i < cnt; i++ {
		dp[i] = make([]bool, cnt)
		dp[i][i] = true
	}
	// 枚举字串长度，最少2，最大字符串长度
	for L := 2; L <= cnt; L++ {
		// 枚举左边界
		for i := 0; i < cnt; i++ {
			// 计算得出右边界
			j := L + i - 1
			// 如果右边界越界，就可以退出当前循环
			if j >= cnt {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			// 只要 dp[i][L] == true 成立，就表示子串 s[i..L] 是回文，此时记录回文长度和起始位置
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

func digui(s string, i, j int) bool {
	if s[i] == s[j] {
		if (j - i) <= 1 {
			return true
		} else if digui(s, i+1, j-1) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
