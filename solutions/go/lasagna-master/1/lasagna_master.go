package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTimePerLayer int) int {
    if prepTimePerLayer == 0 {
        prepTimePerLayer = 2
    }
    return len(layers)*prepTimePerLayer
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodlesGrams int, sauceLiters float64) {
    for _, layer := range layers {
        if layer == "sauce"{
            sauceLiters+=0.2
        } else if layer == "noodles"{
            noodlesGrams+=50
        }
    }
    return
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string){
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(normal []float64, additional int) []float64{
    newScale := make([]float64, len(normal))
    multiple := float64(additional)/2.0
    for i, v := range normal {
        newScale[i] = v * multiple
    }
    return newScale
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
