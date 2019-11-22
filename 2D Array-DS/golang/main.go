package main

import "fmt"

func main() {
	data := [][]int32{{1, 1, 1, 0, 0, 0}, {0, 1, 0, 0, 0, 0}, {1, 1, 1, 0, 0, 0}, {0, 0, 2, 4, 4, 0}, {0, 0, 0, 2, 0, 0}, {0, 0, 1, 2, 4, 0}}
	fmt.Println(hourglassSum(data))
}
func hourglassSum(arr [][]int32) int32 {
	var cons int32 = -1000
	for i := 0; i < 4; i++ {
		for x := 0; x < 4; x++ {
			top := arr[i][x] + arr[i][x+1] + arr[i][x+2]          // lente
			middle := arr[i+1][x+1]                               // puente
			bottom := arr[i+2][x] + arr[i+2][x+1] + arr[i+2][x+2] // lente
			if !(top+middle+bottom <= cons) {
				cons = top + middle + bottom
			}
		}
	}
	return cons
}
