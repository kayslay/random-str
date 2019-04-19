package main

import (
	"fmt"
	"str"
)

func main() {
	fmt.Println(str.WriteFromFormat("Aads*")) // print random strings following the format
	fmt.Println(str.WriteN(10, 'A'))          // return all uppercase
	fmt.Println(str.WriteN(10, 'a'))          // return all lowercase
	fmt.Println(str.WriteN(10, 'd'))          // return all digits
	fmt.Println(str.WriteN(10, 's'))          // return all symbol
	fmt.Println(str.WriteN(10, '*'))          // return all any
}
