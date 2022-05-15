package Stack

import "fmt"

type Stack []interface{}

//IsEmpty- 스택이 비어있는지 확인하는 함수
func (s*Stack)IsEmpty()bool{
	return len(*s) ==0
}
func (s*Stack)Push(data interface{}){
	*s = append(*s,data)//스택끝에 값을 추가함
	fmt.Printf("%d pushed to stack\n",data)
}
func (s*Stack)Pop()interface{}{
	if s.IsEmpty(){
		fmt.Println("Stack is empty")
		return nil

	}else {
		top := len(*s)-1
		data := (*s)[top]
		*s = (*s)[:top]
		return data
	}
}

func Result(){
	var s Stack
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Printf("%d poped form stact\n",s.Pop())
}