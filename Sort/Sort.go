package Sort

import "fmt"
import "sort"
//sort 공부용

func Study1(){
	//sort.Ints(배열)
// 숫자 오름차순 하는 방법
	var nums = []int{3,5,6,1,0,66} 
	fmt.Println("정렬하기 전 :", nums) 
	sort.Ints(nums) 
	fmt.Println("정렬하기 후 :", nums)
}
func Study2(){
		//문자열 오름차순 하는 방법
		//sort.Strings(문자열배열)
	var str = []string{"안녕", "지니", "감기", "조심", "꽃"} 
	fmt.Println("정렬하기 전 :", str) 
	sort.Strings(str) 
	fmt.Println("정렬하기 후 :", str)
}
func Study3(){
	   //숫자 내림차순 하는 방법
	   //sort.Sort(sort.Reverse(sort.IntSlice(배열)))
	var nums2 = []int{3, 5, 6, 1, 0, 66} 
	fmt.Println("정렬하기 전 :", nums2) 
	sort.Sort(sort.Reverse(sort.IntSlice(nums2))) 
	fmt.Println("정렬하기 후 :", nums2)

}

func Study4(){
	//sort.Sort(sort.Reverse(sort.StringSlice(문자열배열)))
	//문자열 내림차순
	var str = []string{"안녕", "지니", "감기", "조심", "꽃"} 
	fmt.Println("정렬하기 전 :", str) 
	sort.Sort(sort.Reverse(sort.StringSlice(str))) 
	fmt.Println("정렬하기 후 :", str)

}
func Study5(){
	//sort.IntsAreSorted(배열)
	//오름차순으로 정렬이 되어 있는 확인 하는 방법
	var nums3 = []int{3, 5, 6, 1, 0, 66}
	 boolean := sort.IntsAreSorted(nums3) 
	 fmt.Println("결과값 :", boolean) //결과값 : false
	  var nums7 = []int{1,3,4,6,9} 
	  boolean = sort.IntsAreSorted(nums7) 
	  fmt.Println("결과값 :", boolean) // 결과값 : true

}
func Study6(){
	//sort.IntSlice(배열).Swap(바꿀인덱스, 바꿀인덱스)
	//배열 자리 교체 하는 방법
	var nums4 = []int{5, 8, 9, 6, 2, 1} 
	fmt.Println("정렬하기 전 :", nums4) 
	sort.IntSlice(nums4).Swap(1, 2) 
	fmt.Println("정렬하기 후 :", nums4)


}
func Study7(){
	//값을 비교해서 boolean 값을 알려주는 함수
	//sort.IntSlice(배열).Less(비교할인덱스, 비교할인덱스)

	var nums4 = []int{5, 8, 9, 6, 2, 1} 
	boolean := sort.IntSlice(nums4).Less(3, 4) // x[i] < x[j] 값을 알려주는 것 
	fmt.Println("결과값 :", boolean) //결과값 : false


}

func Result() { 
	
	

Study1()
Study2()
Study3()
Study4()
Study5()
Study6()

Study7()
 


	 }
