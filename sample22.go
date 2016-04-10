package main

import "fmt"

func main(){
	var a [3][3]int
	b := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(len(a))
	fmt.Println(len(b))
	fmt.Println(a[0][0])
	fmt.Println(b[2][2])
	a[0][0] = 10;
	b[2][2] = 20;
	fmt.Println(a)
	fmt.Println(b)
	}
