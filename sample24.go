package main

import "fmt"

func main(){
	a := [4]int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4, 5, 6, 7, 8}
	c := a[:]
	d := b[2:6]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	c[0] = 10
	d[0] = 100
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	}
