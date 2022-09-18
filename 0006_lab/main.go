package main

import "fmt"

func main() {
	
//Using a for loop and an if statement
// print out all the numbers between 1 and 50 that divisible by 7.	
//Ex1()

/*
Change the code from the previous exercise and use the continue statement to print 
out all the numbers divisible by 7 between 1 and 50
*/
//Ex2()
/*
Change the code from the previous exercise and use the break statement
 to print out only the first 3 numbers divisible by 7 between 1 and 50.
*/
//Ex3()
/*
Using a for loop, an if statement and the logical and operator
 print out all the numbers between 1 and 500 that divisible both by 7 and 5.
*/
//Ex4()

Ex5()
}

func Ex5(){
	age := 10
switch  {
case age<0 || age>100:
	fmt.Println("Invalid Age")
case age<18:
	fmt.Println("You are minor!")
case age==18:
	fmt.Println("Congratulations! You've just become major!")
 default:
	fmt.Println("You are major!")

}

}

func Ex4(){
	for i:=0;i<=500;i++{
		if i%7==0 && i%5==0{
			fmt.Println(i)
		}
	}
}

func Ex3(){

	count:=0
	for i:=1;i<=50;i++{
		if count==3{
			break
		}
		if i%7!=0{
			continue
		}else{
			fmt.Println(i)
			count++
		}

	}
}


func Ex2(){
	for i:=0;i<=50;i++{
		if i%7!=0{
			continue
		}else{
			fmt.Println(i)
		}
		
		
	}
}

func Ex1(){
	for i := 1; i < 50; i++ {
		if i%7==0{
			fmt.Println(i)
		}
	}
}