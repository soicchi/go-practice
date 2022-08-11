package main

import "fmt"

type Stringfy interface {
	toString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) toString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) toString() string {
	return fmt.Sprintf("Number=%v, Model= %v", c.Number, c.Model)
}

func displayString(f Stringfy) {
	fmt.Println(f.toString())
}

func main() {
	person := &Person{Name: "sample", Age: 33}
	car := &Car{Number: "1234", Model: "4856"}

	displayString(person)
	displayString(car)
}
