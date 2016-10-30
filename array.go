// 配列 : スライス
/*
var a1 int
var a2 int

 */

package main

import (
	"fmt"
)

func main() {
	//var a [5]int // a[0] - a[4]
	//a[2] = 3
	//a[4] = 10
	//fmt.Println(a)

	//b := [3]int{1, 3, 5}
	//b := [...]int{1, 3, 5}
	//fmt.Println(b)
	//fmt.Println(len(b))

	a := [5]int{2, 10, 8, 15, 4}
	s := a[2:4] // [8, 15]
	s[1] = 12
	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

}
