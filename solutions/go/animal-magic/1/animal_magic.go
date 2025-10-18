package chance

import (
    "math/rand"
    "time"
    )

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
    a := 1
    b := 20
    rand.Seed(time.Now().UnixNano())
    n := a + rand.Intn(b-a+1) // a ≤ n ≤ b
	return n
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
    rand.Seed(time.Now().UnixNano())
    min := 0.0
    max := 12.0
    // Generate a random float64 in the range [min, max]
	return min + rand.Float64() * (max - min)
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
    animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
    
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(animals), func(i, j int) {
    	animals[i], animals[j] = animals[j], animals[i]
    })

    return animals
}
