package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/pingcap/failpoint"
)

func Test_demo(t *testing.T) {
	failpoint.Enable("github.com/akka/failpoint-demo/failpoint-name", "return(5)")
	result, err := demo()
	spew.Dump(result)
	spew.Dump(err)
}

func Test_demo2(t *testing.T) {
	failpoint.Enable("github.com/akka/failpoint-demo/failpoint-name", "return(false)")
	result, err := demo2()
	spew.Dump(result)
	spew.Dump(err)
}
