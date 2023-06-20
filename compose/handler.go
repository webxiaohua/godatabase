package main

import (
	"encoding/json"
	"fmt"
	"github.com/zehuamama/tinyrpc/compressor"
)

// 压缩类型
type CompressType uint16
const (
	Raw CompressType = iota
	Gzip
	Snappy
	Zlib
)
// 序列化方式
type Serializer interface {
	Marshal(data interface{})([]byte,error)
	Unmarshal(data []byte, obj interface{})(error)
}
// Json方式
type JsonSerializer struct {}
func (*JsonSerializer) Marshal(data interface{})([]byte,error){
	return json.Marshal(data)
}
func (*JsonSerializer) Unmarshal(data []byte, obj interface{})(error){
	return json.Unmarshal(data,obj)
}

// 配置选项
type options struct {
	compressType CompressType
	serializer Serializer
}

//选项
type Option func(o *options)

// 包装压缩类型
func WithCompress(c CompressType) Option {
	return func(o *options) {
		o.compressType = c
	}
}
// 包装序列化方式
func WithSerializer(s Serializer) Option{
	return func(o *options) {
		o.serializer = s
	}
}

func NewServer(opts ...Option){

}

func main()  {
	fmt.Println("hello")

}

