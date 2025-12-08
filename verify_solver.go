package main

import "fmt"

func main() {
	fmt.Println("=== RUBIK'S CUBE MOVE REVERSAL TEST ===\n")

	// Create solved cube
	cube := NewCube()
	fmt.Println("Starting with solved cube")
	fmt.Println("State:", cube.toKociembaString())
	fmt.Println()

	// Perform scramble: R U R' U R U2 R'
	fmt.Println("Applying scramble: R U R' U R U2 R'")
	scramble := []Move{R, U, Ri, U, R, U, U, Ri}
	for _, move := range scramble {
		cube.ApplyMove(move)
	}
	fmt.Println("Scrambled state:", cube.toKociembaString())
	fmt.Println()

	// Create model with move history
	m := &model{
		cube:        cube,
		moveHistory: scramble,
	}

	// Solve using move reversal
	fmt.Println("Solving with move reversal...")
	solution := m.solveCube()
	fmt.Printf("Solution: %d moves\n", len(solution))
	fmt.Print("Moves: ")
	for _, move := range solution {
		fmt.Print(move, " ")
	}
	fmt.Println("\n")

	// Apply solution
	fmt.Println("Applying solution...")
	for _, move := range solution {
		cube.ApplyMove(move)
	}

	// Check result
	finalState := cube.toKociembaString()
	fmt.Println("Final state:", finalState)
	fmt.Println("Expected:    UUUUUUUUURRRRRRRRRFFFFFFFFFDDDDDDDDDLLLLLLLLLBBBBBBBBB")

	if cube.IsSolved() {
		fmt.Println("\n✅ SUCCESS: Cube solved!")
	} else {
		fmt.Println("\n❌ FAILED: Cube not solved")
	}
}
