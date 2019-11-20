package main

import "fmt"

func main() {
	Alice := []int32{1,34, 234}
	Bob := []int32{12,35, 34}
	fmt.Println(compareTriplets(Alice,Bob))
}
func compareTriplets(a []int32, b []int32) []int32 {
	 sum:= []int32{0,0}
	for i := 0; i < 3; i++ {
		value,bool := checker(a[i],b[i])
		if bool {
			sum[value] += 1
		}}
	return sum
}

func checker(i int32, i2 int32)(int32,bool){
	if i ==  i2 {
		return 0,false
	}else{
		if i > i2 {return 0,true} else{ return 1,true}
	}
}