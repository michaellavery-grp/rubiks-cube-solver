package main

import (
	"fmt"
	"github.com/daosyn/kociemba"
)

func main() {
	fmt.Println("=== KOCIEMBA DIRECT TEST ===\n")

	// Test 1: Solved cube
	fmt.Println("Test 1: Solved cube")
	solvedCube := "UUUUUUUUURRRRRRRRRFFFFFFFFFDDDDDDDDDLLLLLLLLLBBBBBBBBB"
	fmt.Println("Input:", solvedCube)
	solution1 := kociemba.Solve(solvedCube)
	fmt.Println("Solution:", solution1)
	fmt.Println()

	// Test 2: Get a random scramble from Kociemba
	fmt.Println("Test 2: Random scramble")
	scramble, layout := kociemba.GetNewScramble()
	fmt.Println("Scramble moves:", scramble)
	fmt.Println("Scrambled layout:", layout)
	solution2 := kociemba.Solve(layout)
	fmt.Println("Solution:", solution2)
	fmt.Println()

	// Test 3: Simple known scramble (just R move)
	fmt.Println("Test 3: After R move")
	// After R move from solved position, the cube state would be:
	// This is a manually calculated state after R
	afterR := "UUFUUFUUFRRRRRRRRRUFFFFFFFF DDDDDDDDDBLLBLLLLL BBBBBBBBB"
	afterR = "UUFUUFUUFRRRRRRRRRFFFFFFFFFDDDDDDDDDBLLBLLLLLBBBBBBBBB" // no spaces
	fmt.Println("Input:", afterR)
	solution3 := kociemba.Solve(afterR)
	fmt.Println("Solution:", solution3)
}
