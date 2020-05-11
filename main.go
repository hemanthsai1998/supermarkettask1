package main

import (
	supermarket "https://github.com/hemanthsai1998"
)

func main() {
	supermarket.Print()
	supermarket.Get("c")
	supermarket.Post("c", 9875)
	supermarket.Post("d", 59.67)
	supermarket.Post("a", 453.5)
	supermarket.Post("b", 00.56)
	supermarket.Print()
	supermarket.Get("c")
	supermarket.Put("a", 45.64)
	supermarket.Delete("b")
	supermarket.Get("b")
	supermarket.Print()
}
