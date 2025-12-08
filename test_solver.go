package main

import (
	"fmt"
)

// Simple test to verify the solver works
func testSolver() {
	fmt.Println("=== RUBIK'S CUBE SOLVER TEST ===\n")

	// Test 1: Verify solved cube format
	cube := NewCube()
	fmt.Println("Test 1: Solved cube state conversion")
	cubeString := cube.toKociembaString()
	fmt.Println("Generated:", cubeString)
	fmt.Println("Expected:  UUUUUUUUURRRRRRRRRFFFFFFFFFDDDDDDDDDLLLLLLLLLBBBBBBBBB")

	if cubeString == "UUUUUUUUURRRRRRRRRFFFFFFFFFDDDDDDDDDLLLLLLLLLBBBBBBBBB" {
		fmt.Println("✅ Solved cube format correct\n")
	} else {
		fmt.Println("❌ Solved cube format incorrect\n")
		return
	}

	// Test 2: Simple scramble (R U R' U')
	fmt.Println("Test 2: Simple scramble (R U R' U')")
	cube.ApplyMove(R)
	cube.ApplyMove(U)
	cube.ApplyMove(Ri)
	cube.ApplyMove(Ui)

	scrambledState := cube.toKociembaString()
	fmt.Println("Scrambled state:", scrambledState)

	// Test 3: Solve it
	fmt.Println("\nTest 3: Solving with Kociemba algorithm...")
	m := &model{
		cube:        cube,
		moveHistory: []Move{R, U, Ri, Ui},
	}

	solution := m.solveCube()
	fmt.Printf("Solution found: %d moves\n", len(solution))

	if len(solution) > 0 {
		fmt.Print("Moves: ")
		for _, move := range solution {
			fmt.Print(move, " ")
		}
		fmt.Println()

		// Apply solution
		for _, move := range solution {
			cube.ApplyMove(move)
		}

		// Check if solved
		finalState := cube.toKociembaString()
		fmt.Println("\nAfter applying solution:", finalState)

		if finalState == "UUUUUUUUURRRRRRRRRFFFFFFFFFDDDDDDDDDLLLLLLLLLBBBBBBBBB" {
			fmt.Println("✅ Cube successfully solved!")
		} else {
			fmt.Println("❌ Cube not solved correctly")
		}
	} else {
		fmt.Println("⚠️  No solution returned (cube may already be solved)")
	}

	fmt.Println("\n=== TEST COMPLETE ===")
	fmt.Println("\nTo run the interactive solver, build and run:")
	fmt.Println("  go build -o rubiks_cube rubiks_cube.go")
	fmt.Println("  ./rubiks_cube")
}

func main() {
	testSolver()
}
