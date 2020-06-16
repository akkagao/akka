package main

import (
	"fmt"
	"log"

	. "github.com/dave/jennifer/jen"
)

// 生成代码，以后生成代码就不用模板了
func main() {
	f := NewFile("main")
	f.Func().Id("main").Params().Block(
		Var().Id("a").Op("=").Lit(10),
		Var().Id("b").Op("=").Lit(20),
		
		Qual("fmt", "Println").Call(Id("demo").Call(Id("a"), Id("b"))),
		Qual("fmt", "Println").Call(Lit("Hello, world")),
	)

	f.Line()

	f.Func().Id("demo").Params(Id("a"), Id("b").Int()).Int().Block(
		Return(Id("a").Op("+").Id("b")),
	)

	fmt.Printf("%#v", f)
	err := f.Save("./result.go")
	if err != nil {
		log.Fatal(err)
	}
}
