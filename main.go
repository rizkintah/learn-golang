package main

import (
	"fmt"
)

func main() {
	cartOne := newCart("User 1")

	// add items
	items := []item{
		{
			itemName: "Flanel",
			qty:      2,
			price:    150000,
		},
		{
			itemName: "Jeans Denim",
			qty:      1,
			price:    500000,
		},
	}

	cartOne.addItems(items...)

	fmt.Println(cartOne)
}
