package DataStructure

import "fmt"

func Array(){
    var a [5]int  // [0 0 0 0 0]
a[4] = 100 // [0 0 0 0 100]

fmt.Println(a,"배열") 
}

func Result(){
    a2 := [...]string{"USA", "China", "India", "Germany", "France"}
    b2 := a2 // a copy of a is assigned to b
    b2[0] = "Singapore"
    fmt.Println("a is ", a2)  // [USA China India Germany France]
    fmt.Println("b is ", b2)  // [Singapore China India Germany France]
    Array()
}