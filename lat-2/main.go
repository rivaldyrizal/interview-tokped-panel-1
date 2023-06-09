package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func reverseNumber(n int) int {

	s1 := strconv.Itoa(n)
	resultString := strings.Split(s1, "")

	sort.Slice(resultString, func(i, j int) bool {
		return resultString[i] > resultString[j]
	})

	value := strings.Join(resultString, "")
	intVar, _ := strconv.Atoi(value)

	return intVar

}

func main() {
	number := -12345

	numberRev := reverseNumber(number)

	fmt.Println(numberRev)

}
