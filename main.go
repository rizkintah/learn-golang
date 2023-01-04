package main

import (
	"fmt"
	"learn-go/library"
)

func main() {
	fmt.Println("Hello World")
	library.UlangData(10)

	var s0 = library.Student{"Rizkinta", 10}
	fmt.Println("Nama akan diubah")
	s0.SetName("Tata")

}
