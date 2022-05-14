package Solution

import "fmt"
import "sort"
//문자열 내림차순으로 배치
func solution3(s string) string { 
	var result string 
	var arr []int 
	for i:=0 ; i<len(s); i++ { 
		arr = append(arr, int(s[i])) 
		} 
		sort.Sort(sort.Reverse(sort.IntSlice(arr))) 
		for i:=0 ; i<len(arr); i++ { 
			result +=string(arr[i]) 
			} 
			return result 
		}

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


	//  func Benchmark_solution(price int, money int, count int) int64 {
	// 	var answer = (count + 1) * count * price / 2 
	// 	fmt.Println("1",price/2) 
	// 	fmt.Println("2",answer) 
	// 	if answer<money { 
	// 		return 0 
	// 		} 
	// 		var excess = int64(answer - money) 
	// 		fmt.Println("5",excess) 
	// 		return excess


	// 	}

		//구구단multiplication table
 func Multiplication (){
	for i:=2 ; i < 10 ; i++ { 
		for j:=2 ; j < 10 ; j++ { 
		fmt.Printf("%d * %d = %d\t", j, i, i * j) 
		} 
	}


 }
 



 //

 //
    func Result() { 
		fmt.Printf(Solution("dada"))
		// fmt.Printf( Benchmark_solution(3,20,4))
		fmt.Printf(solution3("Zbcdefg"))
	
		 }
	