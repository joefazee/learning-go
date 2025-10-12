package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {

	fmt.Println("random number between 1 to 99:", rand.IntN(100))

	fmt.Println("Random IntN(100) [0, 100):")
	for i := 0; i < 5; i++ {
		fmt.Printf("  %d\n", rand.IntN(100)) // Generates an int in [0, 99]
	}
	// Seed is super important (This is currently using the default seed)

	numbers := []int{10, 20, 30, 40, 50}
	fmt.Printf("\nOriginal slice: %v\n", numbers)
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	fmt.Printf("Shuffled slice: %v\n", numbers)

	perm := rand.Perm(5)
	fmt.Printf("\nRandom Permutation of 5 elements: %v\n", perm)

	// The seed problem
	source1 := rand.NewPCG(12345, 67890)
	rng1 := rand.New(source1)
	fmt.Println("Sequence from rng1 (seed 12345, 67890):")
	for i := 0; i < 3; i++ {
		fmt.Printf("  IntN(100): %d, Float64: %.4f\n", rng1.IntN(100), rng1.Float64())
	}

	source2 := rand.NewPCG(12345, 67890) // Same seed
	rng2 := rand.New(source2)
	fmt.Println("\nSequence from rng2 (same seed 12345, 67890):")
	for i := 0; i < 3; i++ {
		fmt.Printf("  IntN(100): %d, Float64: %.4f\n", rng2.IntN(100), rng2.Float64())
	}

	source3 := rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano()+1))
	rng3 := rand.New(source3)
	fmt.Println("\nSequence from rng3:")
	for i := 0; i < 3; i++ {
		fmt.Printf("  IntN(100): %d, Float64: %.4f\n", rng3.IntN(100), rng3.Float64())
	}
}
