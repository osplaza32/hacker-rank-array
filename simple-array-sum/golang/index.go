package main

func simpleArraySum(ar []int32) int32 {
	/*
	 * Write your code here.
	 */
	var sum int32
	for index := range ar {
		// index is the index where we are
		// element is the element from someSlice for where we are
		sum += ar[index]
	}
	return sum
}
func main() {
	letters := []int32{12, 34, 234, 123}
	simpleArraySum(letters)
}
