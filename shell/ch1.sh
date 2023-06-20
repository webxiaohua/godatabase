#!/bin/bash
echo "Hello World"
# 定义变量
name="robin"
echo $name
echo "Hello ${name}"
# 字符串拼接
greeting="Hello "$name"!"
echo $greeting
# 字符串长度
echo "字符串长度："${#greeting}
# 字符串截取 (start:len)
echo ${greeting:1:4}
# 查找子字符串 macos 无效
# string="runoob is a great site"
# echo `expr index "$string" io`  # 输出 4
# 定义数组
arr=(10 2 3)
echo ${arr[@]}
# 数组长度
echo ${#arr[@]}
# 数组单个元素长度
echo ${#arr[0]}