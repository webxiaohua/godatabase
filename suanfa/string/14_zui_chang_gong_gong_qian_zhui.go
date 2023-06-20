// 14.最长公共前缀
package main

import "fmt"

func main(){
	strs:=[]string{"leet","leetcode","leec"}
	res := longestCommonPrefix(strs)
	fmt.Println(res)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}else if len(strs) == 1 {
		return strs[0]
	}else{
		dp := make([]string,len(strs))
		dp[0]=strs[0]
		for i:=1;i<len(strs);i++{
			dp[i] = twoCompare(dp[i-1],strs[i])
			if dp[i] == "" {
				return ""
			}
		}
		return dp[len(strs)-1]
	}
}

func twoCompare(a,b string) (res string) {
	baseStr := b
	compareStr := a
	if len(a) < len(b) {
		baseStr = a
		compareStr = b
	}
	if baseStr == "" {
		return
	}
	for i:=0;i<len(baseStr);i++{
		if baseStr[i] != compareStr[i] {
			res = baseStr[0:i]
			break
		}
		if i+1 == len(baseStr){
			res = baseStr
		}
	}
	return res
}