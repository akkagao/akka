package main

import (
	"encoding/json"
	"fmt"

	"github.com/vmihailenco/msgpack"
)

func main() {
	type Item struct {
		Foo  string
		Name string
		Aget int
	}

	b, err := msgpack.Marshal(&Item{Foo: "bar", Name: "CrazyWolf", Aget: 100})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
	if jb, err := json.Marshal(&Item{Foo: "bar", Name: "CrazyWolf", Aget: 100}); err == nil {
		fmt.Println(string(jb))
	}

	var item Item
	err = msgpack.Unmarshal(b, &item)
	if err != nil {
		panic(err)
	}
	fmt.Println(item.Foo, item.Name, item.Aget)
	// Output: bar
}
