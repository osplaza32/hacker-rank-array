package main

import (
	"fmt"
	"sort"
)

func main() {
	Alice := []int32{1, 34, 234, 34, 234, 234, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 234, 34}
	fmt.Println(birthdayCakeCandles(Alice))

}
func birthdayCakeCandles(ar []int32) int32 {
	sort.Slice(ar, func(i, j int) bool { return ar[i] < ar[j] })
	return int32(FindAndCount(ar, ar[len(ar)-1]))
}
func FindAndCount(a []int32, x int32) int {
	count := 0
	for _, n := range a {
		if x == n {
			count += 1
		}
	}
	return count
}
