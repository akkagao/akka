package main

import (
	"strings"

	"github.com/tal-tech/go-zero/core/stringx"
)

func main() {
	StringxReplace()
	SysReplace()
}

func StringxReplace() {
	replacer := stringx.NewReplacer(map[string]string{
		"__IDFA__": "idfa_value",
		"__IMEI__": "imei_value",
		"__OAID__": "oaid_value",
	})
	replacer.Replace("idfa=__IDFA__&imei=__IMEI__&oaid=__OAID__")
}

func SysReplace() {
	m := map[string]string{
		"__IDFA__": "idfa_value",
		"__IMEI__": "imei_value",
		"__OAID__": "oaid_value",
	}
	str := "idfa=__IDFA__&imei=__IMEI__&oaid=__OAID__"
	for key, value := range m {
		str = strings.Replace(str, key, value, 1)
	}
	// fmt.Println(str)
}
