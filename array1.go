package main

import "fmt"

func main(){
	var a[4]int
	b := [...]int{1, 2, 3, 4, 5, 6, 7,  8}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(len(a))
	fmt.Println(len(b))
	fmt.Println(a[0])
	fmt.Println(b[0])
	a[0] = 10
	b[0] = 20
	fmt.Println(a)
	fmt.Println(b)
}
