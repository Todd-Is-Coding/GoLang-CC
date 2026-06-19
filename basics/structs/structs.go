package main

import "fmt"





type person struct {
	name string
	age int
}

func newPerson (name string) *person {
	p := person {name : name}
	p.age = 21 

	return &p
}




func main(){
	fmt.Println(person{name : "Ahmed" , age: 18})

	fmt.Println(person {name : "Ahmed"})

	fmt.Println(&person {name: "Ahmed"})
	fmt.Println(&person { name: "Ahemd" , age : 18})

	fmt.Println(newPerson("Ahmed"))

	s := person{name : "Ahmed" , age : 21}

	fmt.Println(s.name)

	sp := &s 

	sp.age = 52
	fmt.Println(s.age)
	fmt.Println(sp.age)

	dog :=struct {
		name string 
		isGood bool
	}{
		"Rex",
		true,
	}

	fmt.Println(dog)
	
}