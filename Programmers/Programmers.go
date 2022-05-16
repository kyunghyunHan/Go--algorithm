package Programmers
import "fmt"
func Que1(n int)  {
	//자연수 n이 매개변수로 주어집니다. 
	//n을 x로 나눈 나머지가 1이 되도록 하는 가장 작은 자연수 x를 return 하도록 solution 함수를 완성해주세요.
	// 답이 항상 존재함은 증명될 수 있습니다.

	var x int = 1

    for n%x != 1 {
        x++
    }

    fmt.Println(x)

}
func Que2(n int){
	//정수 n을 입력받아 n의 약수를 모두 더한 값을 리턴하는 함수, solution을 완성해주세요.
	result := 0
    for i := 1; i <= n; i++ {
        if n%i == 0 {
            result += i
        }
    }
	fmt.Println(result)
}

func Result(){
	Que1(12)
	Que2(12)
}