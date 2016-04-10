package main

import "fmt"

func main(){
	a := []int{1, 2, 3, 4}
	for i := 5; i <= 10; i++{
		a = append(a, i)
		fmt.Println(a)
		fmt.Println(len(a))
		fmt.Println(cap(a))
		}
		}
