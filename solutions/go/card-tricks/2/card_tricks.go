package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	favCards := []int{2, 6, 9}
	return favCards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index < 0 || index >= len(slice) {
		return -1
	}
	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index >= 0 && index < len(slice) {
		slice[index] = value
	} else {
		slice = append(slice, value)
	}

	return slice

}

// PrependItems adds an arbitrary number of values at the front of a slice.
// func PrependItems(slice []int, values ...int) []int {
// 	for _, v := range slices.Backward(values) {
// 		slice = append([]int{v}, slice...)
// 	}
// 	return slice
// }

func PrependItems(slice []int, values ...int) []int {
	res := make([]int, 0, len(values)+len(slice))
	res = append(res, values...)
	res = append(res, slice...)
	return res
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	// var newSlice []int
	newSlice := make([]int, 0, len(slice)-1)

	for i, v := range slice {
		if i != index {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}
