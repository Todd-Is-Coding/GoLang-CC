package main

import "fmt"

func plus(a int , b int) int {
	return a+b
}

func plusPlus(a,b,c int) int {
	return a + b + c 
}

func main(){
	a := 1
	b := 2 
	c :=3 
	fmt.Println(plus(a,b))
	fmt.Println(plusPlus(a,b,c))
}