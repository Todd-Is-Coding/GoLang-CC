package main

import "fmt"


func zeroval(i int){
	i = 0
    fmt.Println("zeroval:", i)
}


func zeroptr(i *int){
	*i = 0
	fmt.Println("zeroptr:", *i)
}


func main(){
	i := 1
    fmt.Println("initial:", i)
	zeroval(i)
	zeroptr(&i)
	
	fmt.Println("pointer:", &i)
}