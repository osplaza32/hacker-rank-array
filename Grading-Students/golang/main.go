package main

import "fmt"

func main() {
	letters := []int32{73, 67, 38, 33}
	fmt.Println(gradingStudents(letters))

}
func gradingStudents(grades []int32) []int32 {
	// Write your code here
	var out []int32
	for index := range grades {
		// index is the index where we are
		// element is the element from someSlice for where we are
		if grades[index] >= 38 && grades[index]%5 >= 3 {
			grades[index] = grades[index] + 5 - (grades[index] % 5)
		}
		out = append(out, grades[index])
	}
	return out
}
