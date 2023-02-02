package kata

func Cakes(recipe, available map[string]int) int {
	numOfCakes := 0
	enoughIngredients := true
	for {
		for ingredient, quantity := range recipe {
			if _, ok := available[ingredient]; !ok {
				return 0
			}
			available[ingredient] -= quantity
			if available[ingredient] < 0 {
				enoughIngredients = false
			}
		}
		if enoughIngredients {
			numOfCakes++
		} else {
			break
		}
	}

	return numOfCakes
}
