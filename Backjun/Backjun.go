package Backjun


import (
	"fmt"
	"bufio"
	"os"
)

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

func Back2() {
	//문제
	//전화걸고싶은 번호 있다면 숫자 하나누른다음 금속핀이 있는곳까지 시계방향으로 돌려야한다
	//숫자 하나를 누르면 다이얼이 처음위치로 돌아감
	//숫자 1을 걸려면 총 2초
	//1보다 큰수를 거는데 시간은 이보다 더걸린다
	//한칸옆에 있는 숫자를 걸기위해서는 1초씩 걸린다
	//상근이할머니 는 전화번호를 각 숫자에 해당하는 문자로외운다즉 어떤 단어를 걸때 각 알파벳에 해달하는 숫자를 걸면댄다
	var word string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &word)

	var seconds = []int{3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8, 8, 8, 8, 9, 9, 9, 10, 10, 10, 10}
	var result int
	for i := 0; i < len(word); i++ {
		result += seconds[int(word[i])-65]
	}
	fmt.Println(result,"result")
}
func Result(){
	Back1()
	Back2()
}