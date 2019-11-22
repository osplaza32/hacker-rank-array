package main

import (
	"fmt"
)

func main() {
	letters := []int32{12, 34, 234, 123}
	fmt.Println(rotLeft(letters, -90))

}
func rotLeft(a []int32, d int32) []int32 {
	if int(d) > len(a)-1 {
		d = int32(len(a) - 1)
	}
	return append(a[d:], a[:d]...)
}
