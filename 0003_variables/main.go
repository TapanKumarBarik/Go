package main

import "fmt"

func main() {
	fmt.Println("Integerrs============>")

	var age int = 100

	fmt.Printf("Age = %d\n", age)

	var name string = "Tapan"

	//fmt.Println("name "+name)
	_ = name
	//short had
	i := 1000

	fmt.Printf("%d\n", i*100)

	//declar multiple

	car, cost := "BMW", 10

	fmt.Printf("The cost of the car %v is %d", car, cost)

	opened := false
	_ = opened

	//group

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)
	//swap variables

	k,l:=10,100
	k,l=l,k
	fmt.Printf("k = %d\n l = %d\n",k,l)

	//expression
	sum := 5 + 3.5

	fmt.Println(sum)
}
