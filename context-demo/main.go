package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go doSomeThing(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
	fmt.Println("end")
}
func doSomeThing(ctx context.Context) {
	for {
		time.Sleep(time.Second * 1)
		select {
		case <-ctx.Done():
			fmt.Println("down")
			// 如果调用cancel方法了就退出
			return
		default:
			fmt.Println("working...")
		}
	}
}
