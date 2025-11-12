package lasagna

// TODO: define the 'OvenTime' constant

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
const OvenTime = 40 
func RemainingOvenTime(actual int) int {

return OvenTime - actual
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(layers int) int {
    return 2 * layers
	
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(layers, actual int) int {
	return PreparationTime(layers) + actual
}
