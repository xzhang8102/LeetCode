package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) displayName() {
	fmt.Printf("My name is %s.\n", p.name)
}

func (p *person) setAge(age int) {
	p.age = age
	fmt.Printf("%s is %d years old.\n", p.name, p.age)
}

func main() {
	p := person{
		name: "Bill",
	}
	f1 := p.displayName // f1 execute from its own copy of person
	f1()
	p.name = "Tom"
	f1()
	p = person{
		name: "Bill",
	}
	f2 := p.setAge // f2 execute using the person ptr which affected by changes
	f2(45)
	p.name = "Sammy"
	f2(45)
}
