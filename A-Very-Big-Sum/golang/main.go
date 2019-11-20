package main

import "fmt"

func main() {
	letters := []int64{1000000001,1000000002, 1000000003, 1000000004,1000000005}
	fmt.Println(aVeryBigSum(letters))

}
func aVeryBigSum(ar []int64) int64 {
	var sum int64
	for index := range ar {
		// index is the index where we are
		// element is the element from someSlice for where we are
		sum+= ar[index]
	}
	return sum
}