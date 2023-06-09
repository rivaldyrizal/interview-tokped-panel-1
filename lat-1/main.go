package main

import "fmt"

func findIndex(x int16, arr []int16) []int {
	var r []int
	for i, v := range arr {
		if x == v {
			r = append(r, i)
		}
	}
	return r
}

func main() {
	var arr = []int16{1, 2, 3, 4, 4, 5, 6, 7, 8, 9}
	x := int16(4)

	result := findIndex(x, arr)

	fmt.Println(result)

}
