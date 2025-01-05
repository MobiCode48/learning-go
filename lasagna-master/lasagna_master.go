package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, preparation_time int) int {
	if preparation_time == 0 {
		preparation_time = 2
	}
	return (len(layers)) * preparation_time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	sauce := 0.0
	noodles := 0
	for i := 0; i < len(layers); i++ {
		layer := layers[i]
		if layer == "sauce" {
			sauce++
		} else if layer == "noodles" {
			noodles++
		}
	}
	amountSauce := sauce * 0.2
	amountNoodles := noodles * 50
	return amountNoodles, amountSauce

}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(friendlist []string, mylist []string) {

	last_ingredient_from_friend := friendlist[len(friendlist)-1]
	for i := 0; i < len(mylist); i++ {
		if mylist[i] == "?" {
			mylist[i] = last_ingredient_from_friend
		}
	}

}

// TODO: define the 'ScaleRecipe()' function

func ScaleRecipe(quantities []float64, numberOfPortion int) []float64 {

	scaledQuantities := make([]float64, len(quantities))
	for i := 0; i < len(quantities); i++ {
		elem := quantities[i]
		scaled_elem := elem * (float64(numberOfPortion) / 2)
		scaledQuantities[i] = scaled_elem
	}
	return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
