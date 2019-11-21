package main

import (
	"fmt"
	"time"
)

func main() {
	timea := "07:05:45PM"
	timeb := "07:05:45AM"
	timeConversion(timea)
	timeConversion(timeb)

}
func timeConversion(s string) string {
	/*
	 * Write your code here.
	 */
	t, err := time.Parse("03:04:05PM", s)

	if err != nil {
		fmt.Println(err)
	}
	return t.Format("15:04:05")
}
