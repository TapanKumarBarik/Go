package main

import (
	"fmt"
)



func main(){
    people:=[3]string{"Tapan","kumar","barik"}
 
    friends:=[3]string{"Tapan","kumar","barik"}
    fmt.Println(friends)
    outer:
    for index, name := range people {
        for _, name1 :=range friends{
            if name==name1{
                fmt.Println(index,name)
                break outer
            }
        }
      
       
    }
    fmt.Println("index")
}






// func main() {
// 	fmt.Println("Os Args" ,os.Args)

// 	fmt.Println("Path",os.Args[0] )
// 	fmt.Println("First Argument ", os.Args[1])
// }

// package main

// import (
// 	"fmt"
// )
 
// func main() {
//     for i := 1; i <= 10; i++ {
//         if i%3 == 0 {
//             break 
//         }else {
// 			fmt.Println(i)
// 		}
//     }
// }