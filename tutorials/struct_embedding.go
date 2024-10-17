package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

//container embeds a base
//embedding looks like a field without a name
type container struct {
	base
	str string
}

func main() {

	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	//can access the bases fields directly without having to specify its name
	fmt.Printf("co={num: %v, str: %v}", co.num, co.str)

	fmt.Println("also num:", co.base.num)

	//since container embeds base the methods of base also become methods of a container
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}


	//embedding structs with methods may be used to besto interface implementations onto other structs. a container now implements describer becuase it embeds base
	var d describer = co
	fmt.Println("describer:", d.describe())
}