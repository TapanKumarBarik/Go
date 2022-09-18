package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Os Args" ,os.Args)

	fmt.Println("Path",os.Args[0] )
	fmt.Println("First Argument ", os.Args[1])
}