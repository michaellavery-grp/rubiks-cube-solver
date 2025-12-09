package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Render3DCube creates a 3D perspective ASCII view of the cube
// Based on the Stack Overflow isometric/perspective design
// Shows Top, Left, and Right faces in 3D perspective with proper depth
func (m model) render3DCube() string {
	var lines []string

	// Get cube faces
	// In the image: Top=Blue, Left=Yellow, Right=Red
	top := m.cube.faces[Up]      // Top face
	left := m.cube.faces[Left]   // Left face
	right := m.cube.faces[Front] // Right face (using Front for now)

	border := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("x")
	sp := " "

	// Helper to render a colored character
	cc := func(c Color) string {
		return m.getColorStyle(c).Render(string(m.getColorChar(c)))
	}

	// Helper to render multiple chars
	chars := func(c Color, n int) string {
		s := ""
		for i := 0; i < n; i++ {
			s += cc(c)
		}
		return s
	}

	// Title
	title := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205")).
		Render("Welcome to the rubiks cube!")
	lines = append(lines, title)

	// Top face - isometric diamond view
	// Row 0 of top face (single row at top)
	lines = append(lines, strings.Repeat(sp, 16)+chars(top[0], 1)+chars(top[1], 1)+chars(top[2], 1)+chars(top[2], 1))

	// Row 1
	lines = append(lines, strings.Repeat(sp, 12)+
		chars(top[0], 1)+sp+border+border+
		chars(top[0], 1)+chars(top[1], 1)+chars(top[1], 1)+chars(top[2], 1)+chars(top[2], 1)+chars(top[2], 1)+
		border+border)

	// Row 2
	lines = append(lines, strings.Repeat(sp, 4)+border+
		chars(top[0], 1)+chars(top[0], 1)+chars(top[0], 1)+chars(top[1], 1)+chars(top[1], 1)+chars(top[1], 1)+
		border+border+border+
		chars(top[2], 1)+chars(top[2], 1)+chars(top[2], 1)+chars(top[5], 1)+chars(top[5], 1)+chars(top[5], 1))

	// Row 3 - transition to left and right faces
	lines = append(lines,
		chars(left[0], 1)+chars(left[0], 1)+chars(left[0], 1)+chars(left[0], 1)+
		chars(top[3], 1)+chars(top[3], 1)+border+border+
		chars(top[4], 1)+chars(top[4], 1)+border+border+
		chars(top[5], 1)+chars(top[5], 1)+chars(top[5], 1)+chars(top[5], 1)+border+border+
		chars(right[2], 1)+chars(right[2], 1)+border+border+
		chars(right[5], 1)+chars(right[5], 1)+chars(right[5], 1)+chars(right[5], 1)+sp+border+border+border+border)

	// Now render the left and right faces extending down
	// These faces are 3x3 grids rendered with perspective

	// Left face rows and right face rows (9 rows total, each with 3 stickers)
	for row := 0; row < 3; row++ {
		for subrow := 0; subrow < 3; subrow++ {
			lineNum := row*3 + subrow

			// Calculate indent (decreases as we go down for perspective)
			indent := 0
			if lineNum < 9 {
				indent = 9 - lineNum
				if indent < 0 {
					indent = 0
				}
			}

			line := strings.Repeat(sp, indent)

			// Left face stickers (yellow in example)
			for col := 0; col < 3; col++ {
				idx := row*3 + col
				line += cc(left[idx])
			}

			// Separator
			if subrow == 1 {
				line += border
			} else {
				line += "x"
			}

			// Right face stickers (red in example)
			for col := 0; col < 3; col++ {
				idx := row*3 + col
				line += cc(right[idx])
			}

			// Right border
			if lineNum < 6 {
				line += border + border + border
			} else {
				line += border
			}

			// Add right side chars for depth
			if lineNum < 3 {
				line += chars(right[row*3+2], 1) + chars(right[row*3+2], 1) + chars(right[row*3+2], 1) + chars(right[row*3+2], 1)
			} else if lineNum < 6 {
				line += border
			}

			lines = append(lines, line)
		}
	}

	// Bottom rows (cleanup)
	lines = append(lines, strings.Repeat(sp, 4)+border+chars(left[6], 1)+chars(left[7], 1)+chars(left[7], 1)+
		chars(left[7], 1)+border+border+chars(left[8], 1)+chars(left[8], 1)+border+border+border+border+border+
		chars(right[6], 1)+border+chars(right[7], 1)+chars(right[7], 1)+chars(right[7], 1)+border+
		chars(right[8], 1))

	lines = append(lines, strings.Repeat(sp, 8)+border+chars(left[6], 1)+chars(left[7], 1)+chars(left[7], 1)+
		chars(left[7], 1)+chars(left[7], 1)+chars(left[8], 1)+border+chars(right[6], 1)+chars(right[7], 1)+
		chars(right[7], 1)+chars(right[7], 1)+border+border)

	lines = append(lines, strings.Repeat(sp, 12)+border+chars(left[6], 1)+border+chars(left[8], 1)+chars(left[8], 1)+
		chars(left[8], 1)+chars(left[8], 1)+border+chars(right[6], 1)+border+chars(right[8], 1)+chars(right[8], 1)+border+border)

	lines = append(lines, strings.Repeat(sp, 16)+border+border+chars(left[8], 1)+chars(left[8], 1)+
		chars(left[8], 1)+border+border+chars(right[8], 1)+chars(right[8], 1)+border+border)

	lines = append(lines, strings.Repeat(sp, 20)+border+border+chars(left[8], 1)+chars(left[8], 1)+border+border+
		chars(right[8], 1)+chars(right[8], 1))

	lines = append(lines, strings.Repeat(sp, 24)+border+border+border+border)

	return strings.Join(lines, "\n")
}

// getColorChar returns the character to use for each color
func (m model) getColorChar(c Color) rune {
	// Return space - we'll use background color instead of letters
	return ' '
}

// getColoredBlock returns a colored space character (visible as solid color)
func (m model) getColoredBlock(c Color) string {
	style := m.getColorStyle(c)
	return style.Render(" ")
}
