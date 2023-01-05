package main

type item struct {
	itemName string
	qty      int32
	price    float64
}

type cart struct {
	cartName string
	cartItem []item
	total    float64
}

// buar constractor dgn mengembalikan nilai struct cart
func newCart(cartName string) cart {
	return cart{
		cartName: cartName,
		cartItem: []item{},
		total:    0,
	}
}

// tanda * berarti pointer
func (c *cart) addItems(cart ...item) {
	c.cartItem = append(c.cartItem, cart...)
}
