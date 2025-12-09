package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// SolveWithKociemba solves the cube using Python's Kociemba algorithm
// Returns optimal solution (typically ≤20 moves)
func (m *model) SolveWithKociemba() ([]Move, error) {
	// Convert cube to Kociemba format
	cubeString := m.cube.toKociembaString()

	// Get the directory where the executable is located
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("could not determine script directory")
	}
	scriptDir := filepath.Dir(filename)

	// Path to Python script
	pythonScript := filepath.Join(scriptDir, "kociemba_solver.py")
	venvPython := filepath.Join(scriptDir, "venv", "bin", "python3")

	// Call Python solver
	cmd := exec.Command(venvPython, pythonScript, cubeString)
	output, err := cmd.Output()
	if err != nil {
		// Fallback to system python if venv doesn't work
		cmd = exec.Command("python3", pythonScript, cubeString)
		output, err = cmd.Output()
		if err != nil {
			return nil, fmt.Errorf("kociemba solver error: %v", err)
		}
	}

	// Parse solution string
	solutionString := strings.TrimSpace(string(output))

	// Check for error
	if strings.HasPrefix(solutionString, "ERROR:") {
		return nil, fmt.Errorf(solutionString)
	}

	// Parse moves (handles 2 suffix for 180-degree turns)
	solution := parseKociembaMoves(solutionString)

	return solution, nil
}

// parseKociembaMoves parses Kociemba output including 2 suffix for 180-degree turns
func parseKociembaMoves(solution string) []Move {
	moves := []Move{}
	tokens := strings.Fields(solution)

	for _, token := range tokens {
		// Handle moves with 2 (180-degree turns)
		if strings.HasSuffix(token, "2") {
			base := strings.TrimSuffix(token, "2")
			// Apply the move twice
			switch base {
			case "R":
				moves = append(moves, R, R)
			case "L":
				moves = append(moves, L, L)
			case "U":
				moves = append(moves, U, U)
			case "D":
				moves = append(moves, D, D)
			case "F":
				moves = append(moves, F, F)
			case "B":
				moves = append(moves, B, B)
			}
		} else {
			// Handle regular and prime moves
			switch token {
			case "R":
				moves = append(moves, R)
			case "R'":
				moves = append(moves, Ri)
			case "L":
				moves = append(moves, L)
			case "L'":
				moves = append(moves, Li)
			case "U":
				moves = append(moves, U)
			case "U'":
				moves = append(moves, Ui)
			case "D":
				moves = append(moves, D)
			case "D'":
				moves = append(moves, Di)
			case "F":
				moves = append(moves, F)
			case "F'":
				moves = append(moves, Fi)
			case "B":
				moves = append(moves, B)
			case "B'":
				moves = append(moves, Bi)
			}
		}
	}

	return moves
}
