package Programmers
import "fmt"

var a  =[][]int{{1,4},
{3,2}, {4,1},}

var b=[][]int{{3,3},{3,3},}
func Que4(a, b [][]int) {
	ci, cj := len(a), len(b[0])

    c := make([][]int, ci)
    for ci := range c {
        c[ci] = make([]int, cj)
    }

    for i := 0; i < ci; i++ {
        for j := 0; j < cj; j++ {
            for k := 0; k < len(a[0]); k++ {
                c[i][j] += a[i][k] * b[k][j]
            }
        }
    }

   fmt.Println(c)
}

func Result2(){
   


	Que4(a,b)
	
}