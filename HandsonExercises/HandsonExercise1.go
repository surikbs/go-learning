// package main

// import "fmt"

// var x int = 42
// var y string = "James Bond"
// var z bool = true

// func main() {
// x := 42
// y := "James Bond"
// z := true
// fmt.Println(x, y, z)
// fmt.Println(x)
// fmt.Println(y)
// fmt.Println(z)

// s := fmt.Sprintf("%v \t %v \t %v ", x, y, z)

// fmt.Println(s)
// }

// Exercise go 4 and 5

package main

import "fmt"

type hotdog int

var x hotdog
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	//exercise 5
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)

}
