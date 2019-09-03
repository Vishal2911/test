package main

import "fmt"

func main(){
const(
	x = iota
	y = iota
	z= iota
	w 
)
fmt.Println(x,y,z,w)
var arr[5] int
	fmt.Println(arr[1])
	Arraya := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
Slicea := Arraya[2 :5]
fmt.Println(string(Slicea))
fmt.Println(cap(Slicea))
fmt.Println(len(Slicea))
// arr := []int{1, 2, 3}
// tmp := make([]int, len(arr))
// copy(tmp, arr)
// fmt.Println(tmp)
// fmt.Println(arr)

arrdfg := [] int{1, 2, 3}
tmp := make([]int, len(arrdfg))
copy(tmp, arrdfg)
fmt.Println(tmp)
fmt.Println(arrdfg)

rating := map[string]float32 {"C":5, "Go":4.5, "Python":4.5, "C++":2 }
// map has two return values. For the second return value, if the key doesn't
//exist，'ok' returns false. It returns true otherwise.
csharpRating, ok := rating["C#"]
if ok {
    fmt.Println("C# is in the map and its rating is ", csharpRating)
} else {
    fmt.Println("We have no rating associated with C# in the map")
}

fmt.Print("Enter number : ")
var n int
fmt.Scanln(&n)
/*  Conditional Statement if .... else ........     */
if(n%2==0){
	fmt.Println(n,"is Even number")
}else{
	fmt.Println(n,"is Odd number")
}

}