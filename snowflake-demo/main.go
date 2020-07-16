package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bwmarrin/snowflake"
)

var n *snowflake.Node

func main() {
	Node, err := snowflake.NewNode(1)
	n = Node

	if err != nil {
		println(err)
		os.Exit(1)
	}

	for i := 0; i < 3; i++ {
		id := n.Generate()
		fmt.Println("id", id)
		fmt.Println("node: ", id.Node(), "step: ", id.Step(), "time: ", id.Time())
	}
	for i := 0; i < 3; i++ {
		id := n.Generate()
		fmt.Println("id", id)
		fmt.Println("node: ", id.Node(), "step: ", id.Step(), "time: ", id.Time())
	}
	go test()
	time.Sleep(1000000)
}
func test() {
	for i := 0; i < 3; i++ {
		id := n.Generate()
		fmt.Println("--id", id)
		fmt.Println("node: ", id.Node(), "step: ", id.Step(), "time: ", id.Time())
	}
}
