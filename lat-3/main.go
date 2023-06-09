package main

import (
	"fmt"
	"sort"
)

func combineArray(arr1, arr2 []int16) []int16 {

	var result []int16

	result = append(result, arr1...)

	result = append(result, arr2...)

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}

func main() {
	var arr1 = []int16{2, 3, 4, 5, 6, 7}
	var arr2 = []int16{2, 3, 4, 5, 6, 7}

	arrCombine := combineArray(arr1, arr2)

	fmt.Println(arrCombine)

}
