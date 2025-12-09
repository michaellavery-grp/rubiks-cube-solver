package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Render3DCubeColored creates a colored 3D perspective ASCII view
// Uses actual terminal colors (colored spaces) instead of letters
// Based on the Stack Overflow ascii_cube.png design
func (m model) render3DCubeColored() string {
	var lines []string

	// Get cube faces
	top := m.cube.faces[Up]      // Top face (blue in image)
	left := m.cube.faces[Left]   // Left face (yellow in image)
	right := m.cube.faces[Front] // Right face (red in image)

	// Border character
	border := "x"
	sp := " "

	// Helper to create colored block (2 chars wide for better visibility)
	colorBlock := func(c Color) string {
		style := m.getColorStyle(c)
		return style.Render("  ") // Two spaces with colored background
	}

	// Single color block
	cb := func(c Color) string {
		style := m.getColorStyle(c)
		return style.Render(" ")
	}

	// Title
	title := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205")).
		Render("Welcome to the rubiks cube!")
	lines = append(lines, title)

	// Top face - shown in isometric perspective
	// Pattern matches the image exactly

	// Row 1: Top edge of blue face (4 blocks)
	lines = append(lines, strings.Repeat(sp, 16)+
		cb(top[0])+cb(top[1])+cb(top[2])+cb(top[2]))

	// Row 2: Expanding blue
	lines = append(lines, strings.Repeat(sp, 12)+
		cb(top[0])+sp+border+border+
		cb(top[0])+cb(top[1])+cb(top[1])+cb(top[2])+cb(top[2])+cb(top[2])+
		border+border)

	// Row 3: More blue expansion
	lines = append(lines, strings.Repeat(sp, 4)+border+
		cb(top[0])+cb(top[0])+cb(top[0])+cb(top[1])+cb(top[1])+cb(top[1])+
		border+border+border+
		cb(top[2])+cb(top[2])+cb(top[2])+cb(top[5])+cb(top[5])+cb(top[5]))

	// Row 4: Top face complete, start of left/right faces
	lines = append(lines,
		cb(left[0])+cb(left[0])+cb(left[0])+cb(left[0])+
		cb(top[3])+cb(top[3])+border+border+
		cb(top[4])+cb(top[4])+border+border+
		cb(top[5])+cb(top[5])+cb(top[5])+cb(top[5])+border+border+
		cb(right[2])+cb(right[2])+border+border+
		cb(right[5])+cb(right[5])+cb(right[5])+cb(right[5])+sp+border+border+border+border)

	// Middle section - left and right faces cascading down
	// Pattern: yellow on left, blue in middle (fading), red on right

	// Row 5
	lines = append(lines,
		cb(left[0])+cb(left[0])+cb(left[0])+cb(left[0])+cb(left[0])+cb(left[0])+
		border+border+border+
		cb(top[4])+cb(top[4])+cb(top[4])+cb(top[4])+cb(top[4])+cb(top[4])+
		border+border+border+
		border+border+
		cb(right[2])+cb(right[2])+cb(right[2])+cb(right[2])+cb(right[2])+
		border+border+border+
		cb(right[5])+cb(right[5])+cb(right[5])+cb(right[5])+
		border+border+border+border)

	// Row 6
	lines = append(lines,
		cb(left[0])+cb(left[0])+cb(left[0])+cb(left[0])+cb(left[0])+cb(left[0])+
		border+
		cb(left[3])+cb(left[3])+cb(left[3])+cb(left[3])+cb(left[3])+
		border+border+border+
		cb(top[4])+cb(top[4])+cb(top[4])+cb(top[4])+cb(top[4])+cb(top[4])+cb(top[4])+cb(top[4])+cb(top[4])+cb(top[4])+
		border+border+border+border+border+
		cb(right[1])+cb(right[1])+cb(right[1])+cb(right[1])+
		border+border+
		cb(right[5])+cb(right[5])+cb(right[5])+cb(right[5])+cb(right[5])+cb(right[5])+cb(right[5])+cb(right[5]))

	// Row 7
	lines = append(lines,
		cb(left[0])+cb(left[0])+cb(left[0])+cb(left[0])+cb(left[0])+cb(left[0])+
		border+
		cb(left[3])+cb(left[3])+cb(left[3])+cb(left[3])+cb(left[3])+
		border+
		cb(left[3])+cb(left[3])+cb(left[3])+cb(left[3])+
		border+border+border+
		cb(top[4])+cb(top[4])+cb(top[4])+cb(top[4])+cb(top[4])+cb(top[4])+cb(top[4])+cb(top[4])+cb(top[4])+cb(top[4])+cb(top[4])+cb(top[4])+cb(top[4])+cb(top[4])+
		border+border+border+border+border+border+
		cb(right[1])+cb(right[1])+cb(right[1])+cb(right[1])+
		border+border+
		cb(right[4])+cb(right[4])+cb(right[4])+cb(right[4])+cb(right[4])+cb(right[4])+cb(right[4])+cb(right[4]))

	// Continue pattern down - simplified for now
	// This creates the cascading effect seen in the image
	for i := 0; i < 9; i++ {
		indent := 9 - i
		if indent < 0 {
			indent = 0
		}

		line := strings.Repeat(sp, indent)

		// Left face (yellow)
		leftIdx := (i / 3) * 3
		for j := 0; j < 3; j++ {
			if leftIdx+j < 9 {
				line += cb(left[leftIdx+j])
			}
		}

		line += border

		// Right face (red)
		rightIdx := (i / 3) * 3
		for j := 0; j < 3; j++ {
			if rightIdx+j < 9 {
				line += cb(right[rightIdx+j])
			}
		}

		line += border

		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}
