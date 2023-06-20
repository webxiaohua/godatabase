package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

const(
	CommentRemark  = ';' // 注释符号
)
var (
	AsmFileName string // 待执行的汇编码文件名
	Debug bool // 是否打开调试
	Code [][]string // 存放pcode
	Stack []string // 执行栈
	VarTable map[string]interface{} // 定义表
	Eip int // 栈指针

)

// 变量命名是否合法
func IsValidVariable(ident string) bool {
	if ident == "" {
		return false
	}
	if isMatch,_ := regexp.MatchString(`^_*[A-Za-z]+$`,ident); !isMatch {
		return false
	}
	return true
}

// 检测定义的函数
func CheckLabel(label string) bool {
	if len(label) == 0 {
		return false
	}
	// 解析label， 如： (FUNC @sum:)
	arrs := strings.Split(label," @")
	if len(arrs) == 1 {
		// 普通label，用于跳转
		return IsValidVariable(label)
	}else if len(arrs) == 2 {
		// todo 暂时先不处理
		if arrs[0] == "FUNC" {
			return false
		}
	}
	return false
}

// 解析汇编代码
func Assemb(fileName string){
	// 读取文件内容，提取汇编代码
	lines := make([]string,0)
	f,err := os.Open(fileName)
	if err != nil {
		fmt.Println("error:",err)
		return
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for{
		line,err := r.ReadString('\n')
		if err == io.EOF {
			break
		}else if err != nil {
			fmt.Println("error:",err)
			break
		}
		if line[0] == CommentRemark {
			// 去除注释
			continue
		}
		// 去除前后空格、换行符
		line = strings.ReplaceAll(string(bytes.Trim([]byte(line)," ")),"\n","")
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}
	fmt.Println(lines)
	// todo 解析是否为label

}


func init()  {
	flag.StringVar(&AsmFileName,"file","/Users/shenxinhua/workspace/code/go/src/godatabase/bianyiqi/1.asm","汇编代码文件名")
	flag.BoolVar(&Debug,"debug",false,"是否开启调试")
}

func main() {
	flag.Parse()
	if AsmFileName == "" {
		fmt.Println("no file")
		return
	}
	Assemb(AsmFileName)
}