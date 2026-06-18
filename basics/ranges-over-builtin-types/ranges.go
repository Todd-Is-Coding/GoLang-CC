package main

import "fmt"





func main(){

	nums := []int {1,2,3}
	sum := 0
	for _,num := range nums {
		sum += num
	}

	fmt.Println(sum)

	for i , num := range nums {
		fmt.Println(i , num)
	}

	kvs := map[string]string {"A" : "Apple" , "B" : "Banana"}

	for k,v := range kvs {
		fmt.Printf("%s -> %s\n" ,k , v)
	}

	for i,c := range "go"{
		fmt.Println(i, c)
	}

}