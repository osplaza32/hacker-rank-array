package main

import (
	"fmt"
	"math"
)

func main() {
	letters := []int32{12,-34, -234,123,10,0}
	plusMinus(letters)
}
func plusMinus(arr []int32) {
	c := count{countarr:len(arr),countMinus:0,countPlus:0,countZero:0}
	for index := range arr {
		c.CheckerNumb(arr[index])
	}
	c.work()
}
type count struct {
	countarr int
	countZero int
	countPlus int
	countMinus int
}

func (c *count) CheckerNumb(in int32) {
	if in == 0 {
		c.countZero += 1
	}else{
		if !math.Signbit(float64(in)) { c.countPlus += 1 } else{c.countMinus += 1 }

	}
}

func (c *count) work() {
	fmt.Println(format((float64(c.countPlus) / float64(c.countarr))))
		
	fmt.Println(format(float64(c.countMinus) / float64(c.countarr)))
	fmt.Println(format(float64(c.countZero) / float64(c.countarr)))

}

func format(f float64) interface{} {
return  fmt.Sprintf("%.6f", f) // s == "12.35"

}
