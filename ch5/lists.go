package main

import "fmt"

func main() {

    fmt.Println("More complex built-in data structures (arrays,maps slices)")

    var x [10]int

    for i := 0;i < 10; i++ {

        x[i] = i

        fmt.Println(i," -> ",x[i])

    }

}
