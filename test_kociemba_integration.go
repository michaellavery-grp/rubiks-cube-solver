package main

import "fmt"

func main() {
	fmt.Println("=== KOCIEMBA INTEGRATION TEST ===\n")

	// Create a solved cube
	cube := NewCube()
	fmt.Println("Starting with solved cube")
	fmt.Println("State:", cube.toKociembaString())
	fmt.Println()

	// Apply a scramble
	fmt.Println("Applying scramble: R U R' U R U2 R' U")
	scramble := []Move{R, U, Ri, U, R, U, U, Ri, U}
	for _, move := range scramble {
		cube.ApplyMove(move)
	}
	fmt.Println("Scrambled state:", cube.toKociembaString())
	fmt.Println()

	// Create model
	m := &model{
		cube:        cube,
		moveHistory: scramble,
	}

	// Test Kociemba solver
	fmt.Println("Solving with Kociemba algorithm...")
	solution, err := m.SolveWithKociemba()
	if err != nil {
		fmt.Printf("❌ Kociemba error: %v\n", err)
		fmt.Println("\nFalling back to move reversal...")

		// Try move reversal
		solution = []Move{}
		for i := len(m.moveHistory) - 1; i >= 0; i-- {
			solution = append(solution, reverseMove(m.moveHistory[i]))
		}
		fmt.Printf("Move reversal solution: %d moves\n", len(solution))
	} else {
		fmt.Printf("✅ Kociemba solution: %d moves\n", len(solution))
		fmt.Print("Moves: ")
		for _, move := range solution {
			fmt.Print(move, " ")
		}
		fmt.Println()
	}

	// Apply solution
	fmt.Println("\nApplying solution...")
	for _, move := range solution {
		cube.ApplyMove(move)
	}

	// Check result
	finalState := cube.toKociembaString()
	expectedState := "UUUUUUUUURRRRRRRRRFFFFFFFFFDDDDDDDDDLLLLLLLLLBBBBBBBBB"

	fmt.Println("Final state:", finalState)
	fmt.Println("Expected:   ", expectedState)

	if cube.IsSolved() {
		fmt.Println("\n✅ SUCCESS: Cube solved!")
	} else {
		fmt.Println("\n❌ FAILED: Cube not fully solved")
		fmt.Println("Note: Kociemba may have solved to a different configuration")
	}
}
