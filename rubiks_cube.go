// Rubik's Cube - Isometric 3D ASCII solver using Bubble Tea
// Features: 3D rendering, solving algorithms, move hints, custom input
package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/daosyn/kociemba"
)

// Color represents a cube face color
type Color int

const (
	White Color = iota
	Red
	Blue
	Orange
	Green
	Yellow
)

func (c Color) String() string {
	return [...]string{"W", "R", "B", "O", "G", "Y"}[c]
}

// Cube represents the Rubik's Cube state
// Faces: Front, Right, Back, Left, Up, Down
type Cube struct {
	faces [6][9]Color // 6 faces, 9 stickers each
}

// NewCube creates a solved cube
func NewCube() *Cube {
	c := &Cube{}
	// Initialize solved state with proper color mapping
	// Front=Green, Right=Red, Back=Blue, Left=Orange, Up=White, Down=Yellow
	faceColors := []Color{Green, Red, Blue, Orange, White, Yellow}
	for face := 0; face < 6; face++ {
		for sticker := 0; sticker < 9; sticker++ {
			c.faces[face][sticker] = faceColors[face]
		}
	}
	return c
}

// Face indices
const (
	Front = iota
	Right
	Back
	Left
	Up
	Down
)

// Move represents a cube move
type Move string

const (
	R  Move = "R"  // Right clockwise
	Ri Move = "R'" // Right counter-clockwise
	L  Move = "L"
	Li Move = "L'"
	U  Move = "U"
	Ui Move = "U'"
	D  Move = "D"
	Di Move = "D'"
	F  Move = "F"
	Fi Move = "F'"
	B  Move = "B"
	Bi Move = "B'"
)

// ApplyMove performs a move on the cube
func (c *Cube) ApplyMove(m Move) {
	switch m {
	case R:
		c.rotateRight()
	case Ri:
		c.rotateRight()
		c.rotateRight()
		c.rotateRight()
	case L:
		c.rotateLeft()
	case Li:
		c.rotateLeft()
		c.rotateLeft()
		c.rotateLeft()
	case U:
		c.rotateUp()
	case Ui:
		c.rotateUp()
		c.rotateUp()
		c.rotateUp()
	case D:
		c.rotateDown()
	case Di:
		c.rotateDown()
		c.rotateDown()
		c.rotateDown()
	case F:
		c.rotateFront()
	case Fi:
		c.rotateFront()
		c.rotateFront()
		c.rotateFront()
	case B:
		c.rotateBack()
	case Bi:
		c.rotateBack()
		c.rotateBack()
		c.rotateBack()
	}
}

// rotateFace rotates a face 90° clockwise
func (c *Cube) rotateFace(face int) {
	temp := c.faces[face][0]
	c.faces[face][0] = c.faces[face][6]
	c.faces[face][6] = c.faces[face][8]
	c.faces[face][8] = c.faces[face][2]
	c.faces[face][2] = temp

	temp = c.faces[face][1]
	c.faces[face][1] = c.faces[face][3]
	c.faces[face][3] = c.faces[face][7]
	c.faces[face][7] = c.faces[face][5]
	c.faces[face][5] = temp
}

// rotateRight performs R move
func (c *Cube) rotateRight() {
	c.rotateFace(Right)

	temp := [3]Color{c.faces[Front][2], c.faces[Front][5], c.faces[Front][8]}
	c.faces[Front][2] = c.faces[Down][2]
	c.faces[Front][5] = c.faces[Down][5]
	c.faces[Front][8] = c.faces[Down][8]

	c.faces[Down][2] = c.faces[Back][6]
	c.faces[Down][5] = c.faces[Back][3]
	c.faces[Down][8] = c.faces[Back][0]

	c.faces[Back][0] = c.faces[Up][8]
	c.faces[Back][3] = c.faces[Up][5]
	c.faces[Back][6] = c.faces[Up][2]

	c.faces[Up][2] = temp[0]
	c.faces[Up][5] = temp[1]
	c.faces[Up][8] = temp[2]
}

// rotateLeft performs L move
func (c *Cube) rotateLeft() {
	c.rotateFace(Left)

	temp := [3]Color{c.faces[Front][0], c.faces[Front][3], c.faces[Front][6]}
	c.faces[Front][0] = c.faces[Up][0]
	c.faces[Front][3] = c.faces[Up][3]
	c.faces[Front][6] = c.faces[Up][6]

	c.faces[Up][0] = c.faces[Back][8]
	c.faces[Up][3] = c.faces[Back][5]
	c.faces[Up][6] = c.faces[Back][2]

	c.faces[Back][2] = c.faces[Down][6]
	c.faces[Back][5] = c.faces[Down][3]
	c.faces[Back][8] = c.faces[Down][0]

	c.faces[Down][0] = temp[0]
	c.faces[Down][3] = temp[1]
	c.faces[Down][6] = temp[2]
}

// rotateUp performs U move
func (c *Cube) rotateUp() {
	c.rotateFace(Up)

	temp := [3]Color{c.faces[Front][0], c.faces[Front][1], c.faces[Front][2]}
	c.faces[Front][0] = c.faces[Right][0]
	c.faces[Front][1] = c.faces[Right][1]
	c.faces[Front][2] = c.faces[Right][2]

	c.faces[Right][0] = c.faces[Back][0]
	c.faces[Right][1] = c.faces[Back][1]
	c.faces[Right][2] = c.faces[Back][2]

	c.faces[Back][0] = c.faces[Left][0]
	c.faces[Back][1] = c.faces[Left][1]
	c.faces[Back][2] = c.faces[Left][2]

	c.faces[Left][0] = temp[0]
	c.faces[Left][1] = temp[1]
	c.faces[Left][2] = temp[2]
}

// rotateDown performs D move
func (c *Cube) rotateDown() {
	c.rotateFace(Down)

	temp := [3]Color{c.faces[Front][6], c.faces[Front][7], c.faces[Front][8]}
	c.faces[Front][6] = c.faces[Left][6]
	c.faces[Front][7] = c.faces[Left][7]
	c.faces[Front][8] = c.faces[Left][8]

	c.faces[Left][6] = c.faces[Back][6]
	c.faces[Left][7] = c.faces[Back][7]
	c.faces[Left][8] = c.faces[Back][8]

	c.faces[Back][6] = c.faces[Right][6]
	c.faces[Back][7] = c.faces[Right][7]
	c.faces[Back][8] = c.faces[Right][8]

	c.faces[Right][6] = temp[0]
	c.faces[Right][7] = temp[1]
	c.faces[Right][8] = temp[2]
}

// rotateFront performs F move
func (c *Cube) rotateFront() {
	c.rotateFace(Front)

	temp := [3]Color{c.faces[Up][6], c.faces[Up][7], c.faces[Up][8]}
	c.faces[Up][6] = c.faces[Left][8]
	c.faces[Up][7] = c.faces[Left][5]
	c.faces[Up][8] = c.faces[Left][2]

	c.faces[Left][2] = c.faces[Down][0]
	c.faces[Left][5] = c.faces[Down][1]
	c.faces[Left][8] = c.faces[Down][2]

	c.faces[Down][0] = c.faces[Right][6]
	c.faces[Down][1] = c.faces[Right][3]
	c.faces[Down][2] = c.faces[Right][0]

	c.faces[Right][0] = temp[0]
	c.faces[Right][3] = temp[1]
	c.faces[Right][6] = temp[2]
}

// rotateBack performs B move
func (c *Cube) rotateBack() {
	c.rotateFace(Back)

	temp := [3]Color{c.faces[Up][0], c.faces[Up][1], c.faces[Up][2]}
	c.faces[Up][0] = c.faces[Right][2]
	c.faces[Up][1] = c.faces[Right][5]
	c.faces[Up][2] = c.faces[Right][8]

	c.faces[Right][2] = c.faces[Down][8]
	c.faces[Right][5] = c.faces[Down][7]
	c.faces[Right][8] = c.faces[Down][6]

	c.faces[Down][6] = c.faces[Left][0]
	c.faces[Down][7] = c.faces[Left][3]
	c.faces[Down][8] = c.faces[Left][6]

	c.faces[Left][0] = temp[2]
	c.faces[Left][3] = temp[1]
	c.faces[Left][6] = temp[0]
}

// Scramble scrambles the cube with random moves
func (c *Cube) Scramble(moves int) []Move {
	allMoves := []Move{R, Ri, L, Li, U, Ui, D, Di, F, Fi, B, Bi}
	scramble := make([]Move, moves)
	for i := 0; i < moves; i++ {
		scramble[i] = allMoves[i%len(allMoves)]
		c.ApplyMove(scramble[i])
	}
	return scramble
}

// Model for Bubble Tea
type model struct {
	cube        *Cube
	solution    []Move
	currentMove int
	mode        string // "view", "input", "solve"
	inputFace   int
	inputPos    int
	inputColor  Color
	moveHistory []Move
	message     string
}

func initialModel() model {
	cube := NewCube()
	cube.Scramble(20) // Start with scrambled cube

	return model{
		cube:        cube,
		mode:        "view",
		currentMove: 0,
		message:     "Scrambled cube - Press 's' to solve, 'i' to input custom cube, arrows to rotate",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "s":
			// Solve mode
			m.mode = "solve"
			m.solution = m.solveCube()
			m.currentMove = 0
			m.message = fmt.Sprintf("Solution found: %d moves. Press SPACE for next move", len(m.solution))

		case "i":
			// Input mode
			m.mode = "input"
			m.inputFace = 0
			m.inputPos = 0
			m.message = "Input Mode: Use 1-6 for colors (1=W,2=R,3=B,4=O,5=G,6=Y), arrows to navigate"

		case "v":
			m.mode = "view"
			m.message = "View Mode"

		case " ":
			// Next move in solution
			if m.mode == "solve" && m.currentMove < len(m.solution) {
				move := m.solution[m.currentMove]
				m.cube.ApplyMove(move)
				m.moveHistory = append(m.moveHistory, move)
				m.currentMove++
				m.message = fmt.Sprintf("Move %d/%d: %s", m.currentMove, len(m.solution), move)
			}

		case "enter":
			// Undo last move
			if len(m.moveHistory) > 0 {
				lastMove := m.moveHistory[len(m.moveHistory)-1]
				// Apply reverse move
				reverseMove := reverseMove(lastMove)
				m.cube.ApplyMove(reverseMove)
				m.moveHistory = m.moveHistory[:len(m.moveHistory)-1]
				if m.currentMove > 0 {
					m.currentMove--
				}
				m.message = fmt.Sprintf("Undid: %s", lastMove)
			}

		case "r":
			m.cube.ApplyMove(R)
			m.moveHistory = append(m.moveHistory, R)
			m.message = "R"
		case "R":
			m.cube.ApplyMove(Ri)
			m.moveHistory = append(m.moveHistory, Ri)
			m.message = "R'"
		case "l":
			m.cube.ApplyMove(L)
			m.moveHistory = append(m.moveHistory, L)
			m.message = "L"
		case "L":
			m.cube.ApplyMove(Li)
			m.moveHistory = append(m.moveHistory, Li)
			m.message = "L'"
		case "u":
			m.cube.ApplyMove(U)
			m.moveHistory = append(m.moveHistory, U)
			m.message = "U"
		case "U":
			m.cube.ApplyMove(Ui)
			m.moveHistory = append(m.moveHistory, Ui)
			m.message = "U'"
		case "d":
			m.cube.ApplyMove(D)
			m.moveHistory = append(m.moveHistory, D)
			m.message = "D"
		case "D":
			m.cube.ApplyMove(Di)
			m.moveHistory = append(m.moveHistory, Di)
			m.message = "D'"
		case "f":
			m.cube.ApplyMove(F)
			m.moveHistory = append(m.moveHistory, F)
			m.message = "F"
		case "F":
			m.cube.ApplyMove(Fi)
			m.moveHistory = append(m.moveHistory, Fi)
			m.message = "F'"
		case "b":
			m.cube.ApplyMove(B)
			m.moveHistory = append(m.moveHistory, B)
			m.message = "B"
		case "B":
			m.cube.ApplyMove(Bi)
			m.moveHistory = append(m.moveHistory, Bi)
			m.message = "B'"

		// Input mode controls
		case "1", "2", "3", "4", "5", "6":
			if m.mode == "input" {
				colorNum := int(msg.String()[0] - '0')
				m.cube.faces[m.inputFace][m.inputPos] = Color(colorNum - 1)
				m.inputPos++
				if m.inputPos >= 9 {
					m.inputPos = 0
					m.inputFace++
					if m.inputFace >= 6 {
						m.inputFace = 0
						m.mode = "view"
						m.message = "Input complete! Press 's' to solve"
					}
				}
			}

		case "up":
			if m.mode == "input" {
				if m.inputPos >= 3 {
					m.inputPos -= 3
				}
			}
		case "down":
			if m.mode == "input" {
				if m.inputPos < 6 {
					m.inputPos += 3
				}
			}
		case "left":
			if m.mode == "input" {
				if m.inputPos%3 > 0 {
					m.inputPos--
				}
			}
		case "right":
			if m.mode == "input" {
				if m.inputPos%3 < 2 {
					m.inputPos++
				}
			}
		}
	}

	return m, nil
}

// reverseMove returns the reverse of a move
func reverseMove(m Move) Move {
	switch m {
	case R:
		return Ri
	case Ri:
		return R
	case L:
		return Li
	case Li:
		return L
	case U:
		return Ui
	case Ui:
		return U
	case D:
		return Di
	case Di:
		return D
	case F:
		return Fi
	case Fi:
		return F
	case B:
		return Bi
	case Bi:
		return B
	}
	return m
}

// toKociembaString converts our cube representation to Kociemba format
// Format: 54 characters representing URFDLB faces (9 stickers each)
// Our mapping: Front=Green, Right=Red, Back=Blue, Left=Orange, Up=White, Down=Yellow
// Kociemba order: Up, Right, Front, Down, Left, Back
func (c *Cube) toKociembaString() string {
	var result strings.Builder

	// Map our colors to Kociemba letters
	colorToLetter := func(color Color) byte {
		switch color {
		case White:
			return 'U' // Up
		case Red:
			return 'R' // Right
		case Green:
			return 'F' // Front
		case Yellow:
			return 'D' // Down
		case Orange:
			return 'L' // Left
		case Blue:
			return 'B' // Back
		default:
			return 'U'
		}
	}

	// Kociemba expects: UUUUUUUUU RRRRRRRRR FFFFFFFFF DDDDDDDDD LLLLLLLLL BBBBBBBBB
	// Our faces: Front=0, Right=1, Back=2, Left=3, Up=4, Down=5
	// So map to Kociemba order: Up(4), Right(1), Front(0), Down(5), Left(3), Back(2)
	faceOrder := []int{Up, Right, Front, Down, Left, Back}

	for _, faceIdx := range faceOrder {
		for i := 0; i < 9; i++ {
			result.WriteByte(colorToLetter(c.faces[faceIdx][i]))
		}
	}

	return result.String()
}

// parseMoveString converts Kociemba solution string to our Move slice
func parseMoveString(solution string) []Move {
	moves := []Move{}
	tokens := strings.Fields(solution)

	for _, token := range tokens {
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
		// Handle 180-degree moves
		case "R2":
			moves = append(moves, R, R)
		case "L2":
			moves = append(moves, L, L)
		case "U2":
			moves = append(moves, U, U)
		case "D2":
			moves = append(moves, D, D)
		case "F2":
			moves = append(moves, F, F)
		case "B2":
			moves = append(moves, B, B)
		}
	}

	return moves
}

// solveCube uses Kociemba's algorithm for optimal solving
func (m *model) solveCube() []Move {
	// Convert cube to Kociemba format
	cubeString := m.cube.toKociembaString()

	// Solve using Kociemba algorithm (optimal ≤20 moves)
	solutionString := kociemba.Solve(cubeString)

	// Parse solution string into Move slice
	solution := parseMoveString(solutionString)

	return solution
}

func (m model) View() string {
	var s strings.Builder

	// Title
	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		Render("🧊 RUBIK'S CUBE SOLVER 🧊")
	s.WriteString(title + "\n\n")

	// Render isometric cube
	s.WriteString(m.renderIsometricCube())
	s.WriteString("\n\n")

	// Controls
	controls := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render(
		"[r/R] Right  [l/L] Left  [u/U] Up  [d/D] Down  [f/F] Front  [b/B] Back\n" +
		"[s] Solve  [i] Input Cube  [Space] Next Move  [Enter] Undo  [q] Quit")
	s.WriteString(controls + "\n\n")

	// Status message
	msg := lipgloss.NewStyle().
		Foreground(lipgloss.Color("86")).
		Bold(true).
		Render(m.message)
	s.WriteString(msg + "\n")

	// Move history
	if len(m.moveHistory) > 0 {
		history := "Moves: "
		for _, move := range m.moveHistory {
			history += string(move) + " "
		}
		s.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("14")).Render(history) + "\n")
	}

	// Solution progress
	if m.mode == "solve" && len(m.solution) > 0 {
		progress := fmt.Sprintf("Solution Progress: %d/%d", m.currentMove, len(m.solution))
		s.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Render(progress) + "\n")
	}

	return s.String()
}

// renderIsometricCube renders the cube in isometric 3D
func (m model) renderIsometricCube() string {
	var s strings.Builder

	// Render top face (Up)
	s.WriteString("        " + m.renderFaceRow(Up, 0) + "\n")
	s.WriteString("        " + m.renderFaceRow(Up, 1) + "\n")
	s.WriteString("        " + m.renderFaceRow(Up, 2) + "\n")
	s.WriteString("\n")

	// Render middle three faces (Left, Front, Right)
	for row := 0; row < 3; row++ {
		s.WriteString(m.renderFaceRow(Left, row) + " ")
		s.WriteString(m.renderFaceRow(Front, row) + " ")
		s.WriteString(m.renderFaceRow(Right, row) + "\n")
	}
	s.WriteString("\n")

	// Render bottom face (Down)
	s.WriteString("        " + m.renderFaceRow(Down, 0) + "\n")
	s.WriteString("        " + m.renderFaceRow(Down, 1) + "\n")
	s.WriteString("        " + m.renderFaceRow(Down, 2) + "\n")

	return s.String()
}

// renderFaceRow renders a single row of a face
func (m model) renderFaceRow(face int, row int) string {
	start := row * 3
	colors := m.cube.faces[face][start : start+3]

	var result string
	for i, color := range colors {
		pos := start + i
		style := m.getColorStyle(color)

		// Highlight current position in input mode
		if m.mode == "input" && face == m.inputFace && pos == m.inputPos {
			style = style.Reverse(true).Bold(true)
		}

		result += style.Render(" " + color.String() + " ")
	}

	return result
}

// getColorStyle returns the lip gloss style for a color
func (m model) getColorStyle(c Color) lipgloss.Style {
	switch c {
	case White:
		return lipgloss.NewStyle().Background(lipgloss.Color("255")).Foreground(lipgloss.Color("0"))
	case Red:
		return lipgloss.NewStyle().Background(lipgloss.Color("196")).Foreground(lipgloss.Color("255"))
	case Blue:
		return lipgloss.NewStyle().Background(lipgloss.Color("21")).Foreground(lipgloss.Color("255"))
	case Orange:
		return lipgloss.NewStyle().Background(lipgloss.Color("208")).Foreground(lipgloss.Color("0"))
	case Green:
		return lipgloss.NewStyle().Background(lipgloss.Color("28")).Foreground(lipgloss.Color("255"))
	case Yellow:
		return lipgloss.NewStyle().Background(lipgloss.Color("226")).Foreground(lipgloss.Color("0"))
	default:
		return lipgloss.NewStyle()
	}
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
	}
}
