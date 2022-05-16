package Star

import "fmt"
import "math"
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


//행렬의 덧셈
// var arr1  = int{}
func solution(arr1 [][]int, arr2 [][]int) [][]int {
    for i := range arr1 {
        for j := range arr1[i] {
            arr1[i][j] += arr2[i][j]
        }
    }

    return arr1
}


//임의의 양의 정수 n에 대해, n이 어떤 양의 정수 x의 제곱인지 아닌지 판단하려 합니다.
//n이 양의 정수 x의 제곱이라면 x+1의 제곱을 리턴하고, n이 양의 정수 x의 제곱이 아니라면 -1을 리턴하는 함수를 완성하세요.
func Sol2(n int64){
    sqrt := int64(math.Sqrt(float64(n)))
    if sqrt * sqrt == n {
        sqrt++
        fmt.Println(sqrt * sqrt)
    }else{
 fmt.Println(-1)
    }
   
}
func Result(){
    
   Sol2(121)
}