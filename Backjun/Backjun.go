package Backjun

import "fmt"

func Back1(){
	//첫째 줄에 N과 K가 주어진다. (1 ≤ N ≤ 10, 1 ≤ K ≤ 100,000,000)
//둘째 줄부터 N개의 줄에 동전의 가치 Ai가 오름차순으로 주어진다. (1 ≤ Ai ≤ 1,000,000, A1 = 1, i ≥ 2인 경우에 Ai는 Ai-1의 배수)
	
var n, k int
fmt.Scanf("%d %d\n", &n, &k)

a := make([]int, n)
for i := 0; i < n; i++ {
	fmt.Scanf("%d\n", &a[i])
}

var cnt int = 0
for i := n - 1; i >= 0; i-- {
	cnt += k / a[i]
	k %= a[i]
	if k == 0 {
		break
	}
}
fmt.Println(cnt)

}
func Result(){
	Back1()
}