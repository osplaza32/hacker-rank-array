package main

import (
	"fmt"
	"sort"
)

func main() {
	letters := []int32{256741038, 623958417, 467905213, 714532089, 938071625}

	miniMaxSum(letters)
}
func miniMaxSum(arr []int32) {
	//var min int64
	//var max int64
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	fmt.Println(arr)
	fmt.Println(fmt.Sprintf("%[1]d %[2]d\n", sun(arr[:len(arr)-1]), sun(arr[1:])))

}

func sun(ar []int32) int64 {

	var sum int64
	for index := range ar {
		// index is the index where we are
		// element is the element from someSlice for where we are
		sum += int64(ar[index])
	}
	return sum
}
