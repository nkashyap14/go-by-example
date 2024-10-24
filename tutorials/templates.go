package main

import (
	"os"
	"text/template"
)

//there is built in support for creating dynamic content or showing customized output to the user with text/template package
//there is also a sibling package called html/template for same api but additional security features + used for generating html

func main() {

	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")

	if err != nil {
		panic(err)
	}

	//this will also panic if parse returns an error
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	//execute the template so we generate its text with specific value for its actions
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string {
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	Create := func(name, t string) *template.Template{
		return template.Must(template.New(name).Parse(t))
	}

	//ifi data is a struct we can use .FieldName action to access its felds
	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	//same way we accessed for struct we access for map
	t2.Execute(os.Stdout, map[string]string {
		"Name": "Mickey Mouse",
	})

	//if else provides conditional execution for templates. a value is false if its the default value for a type, ie 0, "", nil, etc
	//- trims whitespace
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")

	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	//range block lets us loop trhoguh slices, arrays, maps, or channels
	//{{.}} inside the range block is set to the current item of the iteration
	t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string {
		"Go",
		"Rust",
		"C++",
		"C#",
	})

}