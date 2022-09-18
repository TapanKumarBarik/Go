package main

/*
Coding Exercise #1

Using the var keyword, declare 4 variables called a, b, c, d of type int, float64, bool and string.

Using short declaration syntax declare and assign these values to variables x, y and z:

- 20

- 15.5

- "Gopher!"

Using fmt.Println() print out the values of a, b, c, d, x, y and z.

Try to run the program without error.
*/
import "fmt"

func main() {
	fmt.Println("Hello")
	a:=-20
	b:=-15.5
	c:="Tapan"
	d:=true

	fmt.Println("a = ",a, "b = ", b, "c = ", c," d = ", d)
}