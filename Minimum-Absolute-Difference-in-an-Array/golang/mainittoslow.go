package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	letters := []int32{1, -3, 71, 68, 17}
	fmt.Println(minimumAbsoluteDifference(letters))

}
func minimumAbsoluteDifference(arr []int32) int32 {
	c := control{arrcount: []int32{}, arrorigin: arr}
	sort.SliceStable(arr, c.fill)
	return c.arrcount[0]
}
func filter(a []int32) []int32 {
	n := 0
	for _, x := range a {
		if !keep(x) {
			a[n] = x
			n++
		}
	}
	return a[:n]
}
func keep(i int32) bool {
	return math.Signbit(float64(i))
}

type control struct {
	arrcount  []int32
	arrorigin []int32
}

func (c *control) fill(i, j int) bool {
	c.arrcount = append(c.arrcount, c.arrorigin[j]-c.arrorigin[i])
	c.arrcount = filter(c.arrcount)
	sort.SliceStable(c.arrcount, func(i, j int) bool { return c.arrcount[i] < c.arrcount[j] })
	return c.arrorigin[i] < c.arrorigin[j]
}
