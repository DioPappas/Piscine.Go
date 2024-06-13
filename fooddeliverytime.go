package piscine

// Define a struct for food items
type food struct {
	preptime int
}

// Map to store menu items and their preparation times
var menu = map[string]food{
	"burger":  {preptime: 15},
	"chips":   {preptime: 10},
	"nuggets": {preptime: 12},
}

// FoodDeliveryTime function to get preparation time for an item
func FoodDeliveryTime(order string) int {
	if item, ok := menu[order]; ok {
		return item.preptime
	}
	// Return 404 if item not found in menu
	return 404
}
