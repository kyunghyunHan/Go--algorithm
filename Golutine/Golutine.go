package Golutine

import "fmt"
import "time"

func One(str string){
	for i := 0; i<3; i++{
		time.Sleep(50* time.Millisecond)
		fmt.Println(str)
	}
}
func Two(str string){
	for i := 0; i<3; i++{
		time.Sleep(70* time.Millisecond)
		fmt.Println(str)
	}
}
func Result(){
	go One("welcome")

	//익명함수로 고루틴 생성
	// go func(){
	// 	for i =0; i<3; i++{
	// 		fmt.Println("to")
	// 	}
	// }
	Two("yoongrammer")
}