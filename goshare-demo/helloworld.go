package main
import (
	"fmt"
	"github.com/mineralres/goshare"
  )
  
  func main() {
	
	var s goshare.DataSource
	symbol := pb.Symbol{Exchange: pb.ExchangeType_SSE, Code: "601398"}
	// 获取最新行情
	data, err := s.GetLastTick(&symbol)
	if err != nil {
	  panic(err)
	}
	fmt.Println(data)
  }