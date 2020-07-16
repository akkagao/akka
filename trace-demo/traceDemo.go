package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	fmt.Println("=======以上为模版代码直接使用就好，下面随便写点东西作为测试使用========")

	count := 0
	for i := 0; i < 10000; i++ {
		count += i
	}
	fmt.Println(count)
}
