package main

import "fmt"

func main() {
	res := addTwoNumbers(genList([]int{2,4,3}), genList([]int{5,6,4}))
	fmt.Println(res.toSlice())
}
