package main

import (
	"fmt"

	str "github.com/kayslay/random-str"
)

func main() {
	fmt.Println(str.WriteFromFormat("Aads*")) // print random strings following the format
	fmt.Println(str.WriteN(100, 'A'))         // return all uppercase
	fmt.Println(str.WriteN(100, 'a'))         // return all lowercase
	fmt.Println(str.WriteN(100, 'd'))         // return all digits
	fmt.Println(str.WriteN(100, 's'))         // return all symbol
	fmt.Println(str.WriteN(100, '*'))         // return all any
}
