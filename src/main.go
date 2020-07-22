package main

func main() {
	var a int = 10
	println(a) // value

	var pa *int = &a
	println(pa) // value

	var paValue = *pa
	println(paValue)
}
