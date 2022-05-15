package DataStructure

import "fmt"

func Array(){
    var a [5]int  // [0 0 0 0 0]
a[4] = 100 // [0 0 0 0 100]

fmt.Println(a,"배열") 
}
func Silice(){
    a4 := [5]int{76,77,78,79,80}
    var b4 []int = a4[1:4] //crates s slice from a[1] to a[3]
    b4[0] = 100
    fmt.Println(a4)
    //make함수 사용
    i := make([]int,6,6)

    fmt.Println(i)
    //append를 사용
    cars := []string{ "Ferrari","Honda","Ford"}
    cars = append(cars,"Toyota") 
    fmt.Println(cars)

    //append를 이용하여 슬라이스 뒤에 슬라이스 추가 ( 추가될 슬라이스 변수에 ... 를 추가한다) 
    vaggis := []string{"potatoses","tomatoes","brinjal"}
    fruits := []string{"oranges","apples"}
    food := append(vaggis,fruits...)
    fmt.Println(food)
    //슬라이스의 일부분을 복사하는데 make로 만든 배열에 copy를 이용하여 복사하면 깊은 복사가 된다.  

    countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
    neededCountries := countries[:len(countries)-2]
    countriesCpy := make([]string, len(neededCountries))
    countriesCpy2 := countries[:len(countries)-2]
    copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
    countriesCpy[0] = "korea"
    countriesCpy2[1] = "japan"
    
    fmt.Println(countriesCpy) // 깊은복사 [korea Singapore Germany]
    fmt.Println(countriesCpy2) // 얖은복사 [USA japan Germany]
    fmt.Println(countries) // [USA japan Germany India Australia]
    슬라이스의 일부분을 복사하는데 make로 만든 배열에 copy를 이용하여 복사하면 깊은 복사가 된다.
}
func Result(){
    a2 := [...]string{"USA", "China", "India", "Germany", "France"}
    b2 := a2 // a copy of a is assigned to b
    b2[0] = "Singapore"
    fmt.Println("a is ", a2)  // [USA China India Germany France]
    fmt.Println("b is ", b2)  // [Singapore China India Germany France]
    Array()
    Silice()
}