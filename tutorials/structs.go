package main 

import "fmt"

type person struct {
	name string
	age int
}

//go is a garbage collected language so the returned pointer is safe and will be automatically cleaned up when there are no active references to it

func newPerson(name string) *person {

    p := person{name: name}
    p.age = 42
    return &p
}

func main() {

	fmt.Println(person{"Bob", 20})

	//can explicitly name field while initializing a struct
	fmt.Println(person{name: "Alice", age: 30})

	//ommitted fields are valued to 0
	fmt.Println(person{name: "Fred"})

	//yielding a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	//utilizing a constructor function is typical practice
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	//when using dots with struct pointers the pointers are automatically dereferenced
	fmt.Println(sp.age)

	//structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age)

	//if a single value struct we don't have to give it a name. this is an anonymous type struct. used for table driven tests
	dog := struct {
		name string
		isGood bool
	}{
		"Rex",
		true,
	}

	fmt.Println(dog)
}