package main

import "fmt"

func main() {
	staircase(6)
}
func staircase(n int32) {
	for i:=0;i<int(n);i++{
		for j:=0;j<int(n)-i-1;j++{
			fmt.Print(" ")
		}
		for j:=0;j<=i;j++{
			fmt.Print("#")
		}
		fmt.Println()
	}

}