package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := make(map[string]int)

	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	emptyBill := make(map[string]int)
	return emptyBill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, isUnitOK := units[unit]

	if !isUnitOK {
		return false
	}

	// itemCount := bill[item]

	// if itemCount > 0 {
	// 	bill[item] = bill[item] + unitValue
	// } else {
	// 	bill[item] = unitValue
	// }
	bill[item] += unitValue

	return true

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitValue, isUnitOK := units[unit]

	if !isUnitOK {
		return false
	}

	itemCount, itemInBill := bill[item]

	if !itemInBill {
		return false
	}

	newCount := itemCount - unitValue

	if newCount < 0 {
		return false
	}

	if newCount == 0 {
		delete(bill, item)
		return true
	}

	bill[item] = newCount

	return true

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, ok := bill[item]
	return qty, ok
}
