package pg

func (d *db) GetPrices() map[int]int {
	newMap := make(map[int]int)
	for key, value := range prices {
		newMap[key] = value
	}

	return newMap
}


var prices = map[int]int{ 
    1: 10,
    2: 20,
    3: 30,
    4: 40,
	5: 50,
}
