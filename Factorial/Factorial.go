package Factorial
//N 의 단계 곱 하기
import "fmt"
func Factorial(n uint64)(result uint64) {
    if (n > 0) {
        result = n * Factorial(n-1)
        return result
    }
    return 1
}
func Fac2(){

	var i int = 16
    fmt.Printf("%d      %d", i, Factorial(uint64(i)))
}