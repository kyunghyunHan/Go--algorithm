package Star

import "fmt"

func Star(){
	var a, b int
    fmt.Scan(&a, &b)
    for i:=0 ; i<b; i++ {
        for j:=0 ; j<a; j++ {
            fmt.Printf("*")
        }
        fmt.Print("\n")
    }


}

func Star2() {
    var x , n int
    fmt.Scan(&x,&n)
    var result []int64
    for i := 0 ; i < n; i++ {
        result = append(result, int64(x+x*i))
    }
    fmt.Println(result) 

}