package main

import "fmt"

//Basic Control Structures

func main() {

    fmt.Println("")
    fmt.Println("form 1 of the for loop ..")
    fmt.Println("")
    
    //Initialisation, iteratiin and loop condition all over the place
    i := 1
    for i <= 10 {
        fmt.Println("i am -> ",i)
        i = i + 1
    }

    fmt.Println("")
    fmt.Println("form 2 of the for loop ..")
    fmt.Println("")

    //Initialisation, iteration and loop condition in one expression
    for i:=1; i <=10; i++ {
        fmt.Println("you are -> ", i)
    }

    fmt.Println("")
    fmt.Println("Loop with if conditionals ...")
    fmt.Println("")

    //Including an if statement
    for i:=1; i <=10; i++ {

        if i % 3 == 0 {
         fmt.Println("keep -> ", i)
        }else{
         fmt.Println("skip -> ", i)
        }

    }
    fmt.Println("")

    //Switch statements
    fmt.Println("Switch case with loop ...")
    fmt.Println("")

    for i:=0; i <=10; i++ {
     switch i {
        case 0: fmt.Println("zero")
        case 1: fmt.Println("one")
        case 2: fmt.Println("two")
        case 3: fmt.Println("three")
        case 4: fmt.Println("four")
        case 5: fmt.Println("five")
        case 6: fmt.Println("six")
        case 7: fmt.Println("seven")
        case 8: fmt.Println("eight")
        case 9: fmt.Println("nine")
        case 10: fmt.Println("ten")
        default: fmt.Println("Unknown")
     }
    }

}


