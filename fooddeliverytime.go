package piscine

type food struct {
	preptime int
}

var menu = map[string]food{
	"burger":  {preptime: 15},
	"chips":   {preptime: 10},
	"nuggets": {preptime: 12},
}

func FoodDeliveryTime(order string) int {
	if item, ok := menu[order]; ok {
		return item.preptime
	}

	return 404
}
