package main

import (
	"fmt"
	processbar "github.com/schollz/progressbar/v3"
	"io"
	"net/http"
	"time"
)

func main(){
	req, _ := http.NewRequest("GET", "https://dl.google.com/go/go1.14.2.src.tar.gz", nil)
	resp,_ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	//f, _ := os.OpenFile("go1.14.2.src.tar.gz", os.O_CREATE|os.O_WRONLY, 0644)
	//defer f.Close()
	bar := processbar.DefaultBytes(resp.ContentLength,"downloading...")
	io.Copy(bar,resp.Body)
	time.Sleep(time.Second * 10)
	fmt.Println("exit.")
}
