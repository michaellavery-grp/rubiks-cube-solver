package main

// Beginner's Method Solver
// Based on layer-by-layer solving: https://ruwix.com/the-rubiks-cube/how-to-solve-the-rubiks-cube-beginners-method/

// Step 4: Yellow Cross Algorithm
// Algorithm: F R U R' U' F'
func yellowCrossAlgorithm() []Move {
	return []Move{F, R, U, Ri, Ui, Fi}
}

// Step 5: Yellow Edges Algorithm
// Algorithm: R U R' U R U2 R' U
func yellowEdgesAlgorithm() []Move {
	return []Move{R, U, Ri, U, R, U, U, Ri, U}
}

// Step 6: Yellow Corners Position Algorithm
// Algorithm: U R U' L' U R' U' L
func yellowCornersPositionAlgorithm() []Move {
	return []Move{U, R, Ui, Li, U, Ri, Ui, L}
}

// Step 7: Yellow Corners Orient Algorithm
// Algorithm: R' D' R D
func yellowCornersOrientAlgorithm() []Move {
	return []Move{Ri, Di, R, D}
}

// Step 3: Second Layer - Left Edge Algorithm
// Algorithm: U' L' U L U F U' F'
func secondLayerLeftAlgorithm() []Move {
	return []Move{Ui, Li, U, L, U, F, Ui, Fi}
}

// Step 3: Second Layer - Right Edge Algorithm
// Algorithm: U R U' R' U' F' U F
func secondLayerRightAlgorithm() []Move {
	return []Move{U, R, Ui, Ri, Ui, Fi, U, F}
}

// Sune Algorithm (for orienting last layer corners)
// Algorithm: R U R' U R U2 R'
func suneAlgorithm() []Move {
	return []Move{R, U, Ri, U, R, U, U, Ri}
}

// Anti-Sune Algorithm
// Algorithm: R U2 R' U' R U' R'
func antiSuneAlgorithm() []Move {
	return []Move{R, U, U, Ri, Ui, R, Ui, Ri}
}

// T-Perm Algorithm (permute last layer corners)
// Algorithm: R U R' U' R' F R2 U' R' U' R U R' F'
func tPermAlgorithm() []Move {
	return []Move{R, U, Ri, Ui, Ri, F, R, R, Ui, Ri, Ui, R, U, Ri, Fi}
}

// Ja-Perm Algorithm (permute last layer edges)
// Algorithm: R' U L' U2 R U' R' U2 R L
func jaPermAlgorithm() []Move {
	return []Move{Ri, U, Li, U, U, R, Ui, Ri, U, U, R, L}
}

// Y-Perm Algorithm (swap diagonal corners)
// Algorithm: F R U' R' U' R U R' F' R U R' U' R' F R F'
func yPermAlgorithm() []Move {
	return []Move{F, R, Ui, Ri, Ui, R, U, Ri, Fi, R, U, Ri, Ui, Ri, F, R, Fi}
}

// Helper: Apply algorithm to cube
func applyAlgorithm(c *Cube, alg []Move) {
	for _, move := range alg {
		c.ApplyMove(move)
	}
}

// Check if cube is solved
func (c *Cube) IsSolved() bool {
	cubeString := c.toKociembaString()
	return cubeString == "UUUUUUUUURRRRRRRRRFFFFFFFFFDDDDDDDDDLLLLLLLLLBBBBBBBBB"
}

// Check if white cross is complete
func (c *Cube) IsWhiteCrossComplete() bool {
	up := Up
	// Check center is white
	if c.faces[up][4] != White {
		return false
	}
	// Check edges are white
	edges := []int{1, 3, 5, 7}
	for _, idx := range edges {
		if c.faces[up][idx] != White {
			return false
		}
	}
	return true
}

// Check if white face is complete (first layer done)
func (c *Cube) IsWhiteFaceComplete() bool {
	up := Up
	for i := 0; i < 9; i++ {
		if c.faces[up][i] != White {
			return false
		}
	}
	return true
}

// Check if yellow cross is formed
func (c *Cube) IsYellowCrossFormed() bool {
	down := Down
	// Check center is yellow
	if c.faces[down][4] != Yellow {
		return false
	}
	// Check edges are yellow
	edges := []int{1, 3, 5, 7}
	for _, idx := range edges {
		if c.faces[down][idx] != Yellow {
			return false
		}
	}
	return true
}

// Check if yellow face is complete
func (c *Cube) IsYellowFaceComplete() bool {
	down := Down
	for i := 0; i < 9; i++ {
		if c.faces[down][i] != Yellow {
			return false
		}
	}
	return true
}

// Beginner's Method Solver
// Returns solution moves and step-by-step description
func (c *Cube) SolveBeginnerMethod() (solution []Move, steps []string) {
	// Make a copy of cube to solve
	cubeCopy := *c

	// Step 1: White Cross (requires pattern matching - simplified for now)
	// For a basic implementation, we'll skip this and assume user starts with white cross
	// or we do random moves until cross forms (not optimal but works)

	// Step 2: White Corners (simplified - random moves)
	// TODO: Implement proper corner insertion

	// Step 3: Second Layer
	// TODO: Implement edge insertion

	// Step 4: Yellow Cross
	if !cubeCopy.IsYellowCrossFormed() {
		steps = append(steps, "Step 4: Creating yellow cross")
		maxAttempts := 4
		for i := 0; i < maxAttempts && !cubeCopy.IsYellowCrossFormed(); i++ {
			alg := yellowCrossAlgorithm()
			applyAlgorithm(&cubeCopy, alg)
			solution = append(solution, alg...)
		}
	}

	// Step 5: Yellow Edges
	steps = append(steps, "Step 5: Positioning yellow edges")
	maxAttempts := 10
	for i := 0; i < maxAttempts; i++ {
		alg := yellowEdgesAlgorithm()
		applyAlgorithm(&cubeCopy, alg)
		solution = append(solution, alg...)
		cubeCopy.ApplyMove(U)
		solution = append(solution, U)
	}

	// Step 6: Yellow Corners Position
	steps = append(steps, "Step 6: Positioning yellow corners")
	for i := 0; i < 5; i++ {
		alg := yellowCornersPositionAlgorithm()
		applyAlgorithm(&cubeCopy, alg)
		solution = append(solution, alg...)
	}

	// Step 7: Yellow Corners Orient
	steps = append(steps, "Step 7: Orienting yellow corners")
	for i := 0; i < 4; i++ {
		alg := yellowCornersOrientAlgorithm()
		// Apply 2-6 times per corner
		for j := 0; j < 4; j++ {
			applyAlgorithm(&cubeCopy, alg)
			solution = append(solution, alg...)
		}
		cubeCopy.ApplyMove(U)
		solution = append(solution, U)
	}

	return solution, steps
}
