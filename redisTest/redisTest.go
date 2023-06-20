package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

var cfg struct {
	RedisConfig RedisConfig
}
type RedisConfig struct {
	Name         string // cluster name, get cluster info by cluster name
	Proto        string
	Addr         string
	Auth         string
}

func InitRedis(){
	configPath:="/Users/shenxinhua/workspace/code/go/src/godatabase/redisTest/redis.toml"
	//fd,_ := os.Open(configPath)
	//defer fd.Close()

	//content,_ := ioutil.ReadFile(configPath)
	//fmt.Println(string(content))
	_,err := toml.DecodeFile(configPath,&cfg)
	if err !=nil {
		fmt.Println("error:",err)
	}
	fmt.Printf("%+v",cfg)
}

func main(){
	InitRedis()
}