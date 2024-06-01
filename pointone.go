package piscine

import "fmt"

func PointOne(n *int) {
	*n = 1
}

func main() {
	var num int
	PointOne(&num)
	fmt.Println(num)
}
