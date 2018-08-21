package main

import (
	"fmt"
	"strings"
)

var zimushuzu = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

var xuhao = map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7, "H": 8, "I": 9, "J": 10, "K": 11, "L": 12, "M": 13, "N": 14, "O": 15, "P": 16, "Q": 17, "R": 18, "S": 19, "T": 20, "U": 21, "V": 22, "W": 23, "X": 24, "Y": 25, "Z": 26}

func main() {
	next := "A"
	for i := 0; i < 702; i++ {
		next = getNextNumber(next)
		// fmt.Println(next)
	}
	fmt.Println(next)
}
func getNextNumber(str string) string {

	target := make([]string, len(str))
	flag := false
	for i := len(str) - 1; i >= 0; i-- {
		zimu := string(str[i])
		if flag {
			target[i] = zimu
			continue
		}
		nextNum, isZ := nextLetter(zimu)
		target[i] = nextNum
		flag = !isZ

		if i == 0 && isZ {
			return "A" + strings.Join(target, "")
		}
	}

	return strings.Join(target, "")
}

func nextLetter(zimu string) (string, bool) {
	if strings.Compare(zimu, "Z") == 0 {
		return "A", true
	}
	index := xuhao[zimu]
	return string(zimushuzu[index]), false
}
