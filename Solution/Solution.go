package Solution

import "fmt"
//단어 s의 가운데 글자를 반환하는 함수, solution을 만들어 보세요. 단어의 길이가 짝수라면 가운데 두글자를 반환
func Solution(s string) string{ 
	var a string
	//
	length := len(s)
	if length%2 != 0 {
		 a = s[length/2 : length/2+1] 
		 } else { 
			 a = s[length/2-1 : length/2+1] 
			}

			return a
	 }



    func Solution2() { 
		fmt.Printf(Solution("dada"))
		 }
	