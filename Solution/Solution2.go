package Solution

import "fmt"


//분할접근방식
//-> 큰문제를 작은 단위로 나누워서 해결하고 합치면서 큰 문제에 도달

func maxSubArray(nums []int)int{
	right := len(nums)-1
	return findMaxSubArray(nums,0,right)
}
func findMaxSubArray(nums []int, left int, right int)int{

	if left == right{
		return nums[left]
	}else{
    mid := (left+right)/2
    leftMax:= findMaxSubArray(nums,left, mid)
	rightMax:= findMaxSubArray(nums,mid+1, right)
	crossMax:= maxCrossing(nums,left,right,mid)
		return max(leftMax,rightMax,crossMax)
}



}
func maxCrossing(nums[]int,left int, right int, mid int){
	leftSum := -999999
	rightSum:= -999999
	sum:= 0
	for i := mid; i>=left; i--{
		sum +=nums[i]
		leftSum =max(leftSum,sum)
	}
	for i := mid+1; i<=right; i++{
		sum +=nums[i]
		rightSum =max(rightSum,sum)
	}
	return leftSum + rightSum
}
func max(values ...int){
	maxVal :- values[0]
	for _,v:= range values{
		if v > maxVal{
		maxVal = v
		}
	}
	return maxVal
}
func SubResult(){
	problem := []int{-2,1,-3,4,-1,2,1,-5,4}
	answer := maxSubArray(problem)
	fmt.Println(answer)
}