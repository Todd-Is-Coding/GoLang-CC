package main

import "fmt"



func SliceIndex[S ~[] E , E comparable](s S , v E) int {
	for i := range s {
		if (s[i] == v){
			return i 
		}
	}

	return -1
}


type List[T any] struct {
	head , tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val T
}

func (lst *List[T]) Push (v T) {
	if lst.tail == nil{
		lst.head = &element[T]{val : v}
		lst.tail = lst.head

	}else{
		lst.tail.next =  &element[T]{val : v}
		lst.tail = lst.tail.next 
	}
}


func(lst *List[T]) allElement () []T {
	var elements []T 

	for e := lst.head ; e != nil ; e = e.next {
		elements = append(elements, e.val)
	}

	return  elements
}

func main(){

	var s = []string{"foo", "bar", "zoo"}
	fmt.Println("index of zoo:", SliceIndex(s, "zoo"))

	_ = SliceIndex(s, "zoo")
		
    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
    fmt.Println("list:", lst.allElement())
	
}