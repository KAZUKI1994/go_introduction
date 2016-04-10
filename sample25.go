package main

import "fmt"

func main(){
	a := make([]int, 10, 20)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	}
