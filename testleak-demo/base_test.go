package testleak_demo

import (
	"fmt"
	"testing"

	. "github.com/pingcap/check"
	"github.com/pingcap/tidb/util/testleak"
)

var _ = Suite(&testBaseSuite{})

type testBaseSuite struct {
}

func TestT(t *testing.T) {
	//CustomVerboseFlag = true
	TestingT(t)
}

func (s *testBaseSuite) SetUpSuite(c *C) {
	fmt.Println("SetUpSuite")
}
func (s *testBaseSuite) SetUpTest(c *C) {
	fmt.Println("SetUpTest")
	testleak.BeforeTest()
}

func (s *testBaseSuite) Testdemo(c *C) {
	demo(10)
}

func (s *testBaseSuite) TearDownTest(c *C) {
	fmt.Println("TearDownTest")
	testleak.AfterTest(c)()
}

func (s *testBaseSuite) TearDownSuite(c *C) {
	fmt.Println("TearDownSuite")
}
