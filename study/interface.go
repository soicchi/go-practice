package main

import "fmt"


type stringfy interface {
		toString() string
}

type Person struct {
		Name string
		Age int
}

func (p *Person) toString() string {
		return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
		Number string
		Model string
}

func (c *Car) toString() string {
		return fmt.Sprintf("Number=%v, Model= %v", c.Number, c.Model)
}

func main() {
		vs := []stringfy{
				&Person{Name: "tanaka", Age: 44},
				&Car{Number: "123455", Model: "sdf"},
		}

		for _, v := range vs {
				fmt.Println(v.toString())
		}
}