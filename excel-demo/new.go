package main

import "fmt"

var zimushuzu = []string{"Z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func main() {
	for i := 0; i < 17578; i++ {
		result := getNum(i)
		fmt.Println(i, " : ", result)
	}

	// num := 676
	// result := getNum(num)
	// fmt.Println(num, " : ", result)
}

func getNum(num int) string {

	result := make([]string, 0)

	flag := false
	for ; num > 1; num = num / 26 {
		index := num % 26
		if flag {
			fmt.Println("index:", index)

			if index == 0 {
				result = append(result, "Y")
			} else {
				// fmt.Println(num%26, index, zimushuzu[index-1])
				result = append(result, zimushuzu[index-1])
			}

			flag = false
			continue
		}
		// fmt.Println(num%26, index, zimushuzu[index])
		result = append(result, zimushuzu[index])
		flag = index == 0
	}

	resultStr := ""

	for i := len(result) - 1; i >= 0; i-- {
		resultStr = resultStr + result[i]
	}

	// fmt.Println(resultStr)
	return resultStr
}
