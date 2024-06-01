package piscine

func UltimatePointOne(n ***int) {
	***n = 1
}

func main() {
	var a int
	var b *int = &a
	var n **int = &b
	UltimatePointOne(&n)
}
