package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [][]int32{{-1 ,1 ,-7 ,-8},{-10, -8, -5 ,-2},{0, 9, 7, -1},{0, 9, 7, -1}}
	fmt.Println(diagonalDifference(arr))
}
func diagonalDifference(arr [][]int32) int32 {
	ctldir := control{direccion:0,reverse:false,work:false}
	ctlinv := control{direccion:len(arr[0])-1,reverse:true,work:false}
	ctldir.diagsum(arr)
	ctlinv.diagsum(arr)
	total := ctldir.total-ctlinv.total
	if !math.Signbit(float64(total)) { return total} else{ return total * -1}
}
type control struct {
	work bool
	reverse bool
	direccion int
	total int32
}
func (c *control) diagsum(arr [][]int32) {
	for index := range arr {
		c.total += arr[index][c.direccioner()]
	}
}

func (c *control) direccioner()int{
	if !c.work {
		c.work = true
	}else{
		if !c.reverse {
			c.direccion += 1
		}else {c.direccion-=1}
	}
	return c.direccion
}