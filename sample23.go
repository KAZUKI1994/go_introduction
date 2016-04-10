package main

import "fmt"

func main(){
	var a[8]int
	b := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a)
	fmt.Println(b)
	a = b
	fmt.Println(a)
	fmt.Println(b)
	c := b
	fmt.Println(c)
	}
