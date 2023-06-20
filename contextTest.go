package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	rcCtx,cancel := shrinkTimeoutContext(context.Background(),time.Second * 3)
	defer cancel()
	fmt.Println(GetContent(rcCtx))
}

// 控制子调用的超时，整个GetContent耗时2s，GetRecommendContent限制1秒
func shrinkTimeoutContext(ctx context.Context,timeout time.Duration) (context.Context,context.CancelFunc){
	t,ok := ctx.Deadline()
	if !ok {
		return context.WithTimeout(ctx,timeout)
	}
	nt := time.Now()
	if t.Before(nt) {
		return context.WithTimeout(ctx,0)
	}
	return context.WithDeadline(ctx,nt.Add(timeout))
}

func GetRecommendContent(ctx context.Context) string{
	fmt.Println("GetRecommendContent1 :",time.Now().Unix())
	time.Sleep(time.Second * 1)
	fmt.Println("GetRecommendContent2 :",time.Now().Unix())
	return "recommend content"
}

func GetContent(ctx context.Context) string {
	rcCtx,cancel := shrinkTimeoutContext(ctx,time.Second * 2)
	var res string
	go func() {
		res = GetRecommendContent(rcCtx)
		cancel()
	}()
	select {
	case <-rcCtx.Done():
		if rcCtx.Err() == context.DeadlineExceeded {
			fmt.Println(time.Now().Unix())
			return "ten hot news"
		}
	}
	return res
}