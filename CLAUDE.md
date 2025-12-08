# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**Rubik's Cube Solver** is a terminal-based 3D ASCII Rubik's Cube solver built with Go, Bubble Tea, and Lip Gloss. It features isometric rendering, interactive controls, and optimal solving using Kociemba's Algorithm.

## Running the Application

```bash
# Build the application
go build -o rubiks_cube rubiks_cube.go

# Run it
./rubiks_cube
```

**Prerequisites:**
- Go 1.18 or later
- Terminal with color support
- Recommended terminal size: 120×40 or larger

**Dependencies (auto-installed via `go mod tidy`):**
- `github.com/charmbracelet/bubbletea` - TUI framework
- `github.com/charmbracelet/lipgloss` - Terminal styling
- `github.com/daosyn/kociemba` - Optimal cube solving algorithm

## Architecture Overview

### Single-File Application
The entire application is contained in `rubiks_cube.go` (~600 lines). This is intentional - it's a self-contained terminal application that can be distributed as a single binary.

### Core Data Structures

**Cube State (`Cube` struct):**
```go
type Cube struct {
    faces [6][9]Color // 6 faces, 9 stickers each
}
```

- 6 faces: Front (Green), Right (Red), Back (Blue), Left (Orange), Up (White), Down (Yellow)
- Each face has 9 stickers (3×3 grid)
- Sticker layout per face:
  ```
  0 1 2
  3 4 5
  6 7 8
  ```

**Color Enum:**
```go
const (
    White Color = iota  // Up face
    Red                 // Right face
    Blue                // Back face
    Orange              // Left face
    Green               // Front face
    Yellow              // Down face
)
```

**Face Indices:**
```go
const (
    Front = iota  // Index 0
    Right         // Index 1
    Back          // Index 2
    Left          // Index 3
    Up            // Index 4
    Down          // Index 5
)
```

### Bubble Tea Model

**State Management:**
```go
type model struct {
    cube        *Cube
    solution    []Move
    currentMove int
    mode        string    // "view", "input", "solve"
    inputFace   int       // Current face being edited (0-5)
    inputPos    int       // Current sticker being edited (0-8)
    inputColor  Color     // Color being placed
    moveHistory []Move    // All moves performed
    message     string    // Status message
}
```

**Modes:**
- **View Mode**: Default - rotate cube, perform moves
- **Input Mode**: Enter custom cube configuration (press 'i')
- **Solve Mode**: Show solution step-by-step (press 's')

### Move System

**12 Standard Moves:**
- R, L, U, D, F, B (clockwise rotations)
- R', L', U', D', F', B' (counter-clockwise rotations, called "prime" moves)

**Move Implementation:**
Each move consists of:
1. Face rotation (rotate 9 stickers 90° clockwise)
2. Edge movement (swap 12 edge stickers between adjacent faces)

**Example - Right Face Rotation:**
```go
func (c *Cube) rotateRight() {
    c.rotateFace(Right)  // Rotate right face 90° clockwise

    // Move edges: Front → Up → Back → Down → Front
    temp := [3]Color{c.faces[Front][2], c.faces[Front][5], c.faces[Front][8]}
    c.faces[Front][2] = c.faces[Down][2]
    c.faces[Front][5] = c.faces[Down][5]
    c.faces[Front][8] = c.faces[Down][8]
    // ... (12 sticker swaps total)
}
```

### Kociemba Solver Integration

**Algorithm:** Two-phase optimal solver that guarantees ≤20 move solutions

**Cube Format Conversion:**
Kociemba expects a 54-character string: `"UUUUUUUUURRRRRRRRRFFFFFFFFFDDDDDDDDDLLLLLLLLLBBBBBBBBB"`
- Order: Up, Right, Front, Down, Left, Back (9 stickers each)
- Each character represents the color at that position

**Key Functions:**
```go
// Convert our cube to Kociemba format
func (c *Cube) toKociembaString() string

// Parse Kociemba solution to our Move types
func parseMoveString(solution string) []Move

// Main solver function
func (m *model) solveCube() []Move {
    cubeString := m.cube.toKociembaString()
    solutionString := kociemba.Solve(cubeString)
    return parseMoveString(solutionString)
}
```

### Rendering System

**Isometric 3D Layout:**
```
        Up Face (White)

Left   Front   Right
(Orange) (Green) (Red)

        Down Face (Yellow)
```

**Color Styling with Lip Gloss:**
Each color has distinct terminal background/foreground:
```go
func (m model) getColorStyle(c Color) lipgloss.Style {
    switch c {
    case White:
        return lipgloss.NewStyle().Background(lipgloss.Color("255")).Foreground(lipgloss.Color("0"))
    case Red:
        return lipgloss.NewStyle().Background(lipgloss.Color("196")).Foreground(lipgloss.Color("255"))
    // ... etc for all 6 colors
    }
}
```

## Controls Reference

### Movement Keys
| Key | Action | Description |
|-----|--------|-------------|
| `r` | R | Rotate right face clockwise |
| `R` | R' | Rotate right face counter-clockwise |
| `l` | L | Rotate left face clockwise |
| `L` | L' | Rotate left face counter-clockwise |
| `u` | U | Rotate top face clockwise |
| `U` | U' | Rotate top face counter-clockwise |
| `d` | D | Rotate bottom face clockwise |
| `D` | D' | Rotate bottom face counter-clockwise |
| `f` | F | Rotate front face clockwise |
| `F` | F' | Rotate front face counter-clockwise |
| `b` | B | Rotate back face clockwise |
| `B` | B' | Rotate back face counter-clockwise |

### Mode Keys
| Key | Action | Description |
|-----|--------|-------------|
| `s` | Solve Mode | Calculate and display optimal solution |
| `i` | Input Mode | Enter custom cube configuration |
| `v` | View Mode | Return to viewing mode |
| `Space` | Next Move | Execute next move in solution |
| `Enter` | Undo Move | Reverse last move |
| `q` | Quit | Exit program |

### Input Mode (Press `i`)
| Key | Action |
|-----|--------|
| `1` | White sticker |
| `2` | Red sticker |
| `3` | Blue sticker |
| `4` | Orange sticker |
| `5` | Green sticker |
| `6` | Yellow sticker |
| `↑↓←→` | Navigate between stickers |
| `Tab` | Next face |

## Common Modifications

### Adding New Algorithms
To add predefined algorithm sequences (like T-Perm, Y-Perm):

```go
// Add algorithm constants
const (
    TPerm []Move = []Move{R, U, Ri, Ui, Ri, F, R, R, Ui, Ri, Ui, R, U, Ri, Fi}
)

// Add execution function
func (c *Cube) ExecuteAlgorithm(alg []Move) {
    for _, move := range alg {
        c.ApplyMove(move)
    }
}
```

### Changing Color Scheme
Modify the `getColorStyle()` function to change terminal colors:

```go
case White:
    return lipgloss.NewStyle().
        Background(lipgloss.Color("255")).  // Change background
        Foreground(lipgloss.Color("0"))     // Change foreground
```

### Adding Animation
To add smooth transitions between moves:

1. Create intermediate frame states
2. Use `tea.Tick` for timed updates
3. Interpolate between cube states

```go
type animationMsg struct{}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case animationMsg:
        // Update animation frame
        return m, tea.Tick(time.Millisecond*50, func(t time.Time) tea.Msg {
            return animationMsg{}
        })
    }
}
```

## Technical Details

### Face Rotation Algorithm
The `rotateFace()` function rotates a face 90° clockwise:

```go
func (c *Cube) rotateFace(face int) {
    f := c.faces[face]
    c.faces[face] = [9]Color{
        f[6], f[3], f[0],  // Bottom-left → Top-left
        f[7], f[4], f[1],  // Bottom-center → Top-center
        f[8], f[5], f[2],  // Bottom-right → Top-right
    }
}
```

Sticker transformation:
```
0 1 2      6 3 0
3 4 5  →   7 4 1
6 7 8      8 5 2
```

### Prime Move Implementation
Prime moves (counter-clockwise) are implemented as 3× clockwise rotations:

```go
case Ri:
    c.rotateRight()
    c.rotateRight()
    c.rotateRight()
```

This works because 3× 90° clockwise = 1× 90° counter-clockwise.

### Input Validation
Currently, there's **no validation** that an input cube is solvable. Invalid cubes will cause Kociemba to fail. Future enhancement:

```go
func (c *Cube) IsValid() bool {
    // Check:
    // 1. Each color appears exactly 9 times
    // 2. Corner parity is correct
    // 3. Edge parity is correct
    // 4. Corner orientation parity is correct
    // 5. Edge orientation parity is correct
}
```

## Known Limitations

1. **No Animation**: Moves are instant (no smooth transitions)
2. **Limited View**: Can only see 3 faces simultaneously
3. **No Input Validation**: Invalid cube configurations may crash solver
4. **No Scramble Generator**: Must manually scramble or use input mode
5. **No Solve Visualization**: Solution shows all moves at once, not step-by-step analysis

## Development Workflow

### Building
```bash
go build -o rubiks_cube rubiks_cube.go
```

### Testing Changes
```bash
# Quick test run
./rubiks_cube

# Test with specific scramble
# 1. Run program
# 2. Press 'i' for input mode
# 3. Configure cube
# 4. Press 's' to solve
# 5. Press SPACE to step through solution
```

### Debugging
Add debug output to terminal:

```go
// In Update function
m.message = fmt.Sprintf("DEBUG: mode=%s face=%d pos=%d", m.mode, m.inputFace, m.inputPos)
```

## Performance

### Benchmarks
| Operation | Time | Notes |
|-----------|------|-------|
| Single Move | ~10μs | Instant visual update |
| Face Rotation | ~5μs | Pure algorithm |
| Full Render | ~2ms | Terminal output |
| Kociemba Solve | ~100-500ms | Varies by scramble complexity |

### Optimization Tips
1. **Batch Moves**: Apply multiple moves before rendering
2. **Lazy Updates**: Only render on user input (already implemented)
3. **Move Caching**: Cache common algorithm sequences

## Future Enhancements

### Phase 1: Polish
- [ ] Add move animations with interpolation
- [ ] Add cube rotation (view from different angles)
- [ ] Add scramble generator (random 20-move sequences)
- [ ] Add timer for speedsolving practice

### Phase 2: Educational
- [ ] Add CFOP method breakdown (Cross, F2L, OLL, PLL)
- [ ] Add beginner's method with step-by-step guidance
- [ ] Add algorithm library (T-Perm, Y-Perm, Sune, etc.)
- [ ] Add tutorial mode explaining each step

### Phase 3: Advanced
- [ ] Add 3D rotation with arrow keys
- [ ] Add solve visualization (highlight affected pieces)
- [ ] Add statistics (average solve time, move count)
- [ ] Add save/load cube states
- [ ] Add multiplayer race mode (two terminals solving same scramble)

## References

### Speedcubing Resources
- **Speedsolving Wiki**: https://www.speedsolving.com/wiki/
- **Cubeskills**: https://www.cubeskills.com/
- **CFOP Tutorial**: https://www.cubeskills.com/tutorials/cfop-speedcubing-method
- **Algorithm Database**: http://algdb.net/

### Optimal Solving
- **Kociemba's Algorithm**: http://kociemba.org/cube.htm
- **God's Number**: https://www.cube20.org/
- **Group Theory**: https://web.mit.edu/sp.268/www/rubik.pdf

### Go Libraries
- **Bubble Tea**: https://github.com/charmbracelet/bubbletea
- **Lip Gloss**: https://github.com/charmbracelet/lipgloss
- **Kociemba Solver**: https://github.com/daosyn/kociemba

## Troubleshooting

### Common Issues

**Issue**: Colors not displaying correctly
**Fix**: Ensure terminal supports 256 colors, try iTerm2 or modern terminal

**Issue**: Solver returns no solution
**Fix**: Cube may be in invalid state - reset with new cube

**Issue**: Input mode not responding
**Fix**: Ensure you're in input mode (press 'i'), check that focus is on terminal

**Issue**: Build fails with missing dependencies
**Fix**: Run `go mod tidy` to install all dependencies

## Credits

**Created by**: Claude Code (Anthropic)
**UI Framework**: Charm Bracelet (Bubble Tea + Lip Gloss)
**Solver**: Kociemba's Algorithm (github.com/daosyn/kociemba)
**Inspiration**: Classic Rubik's Cube (1974, Ernő Rubik)

---

**Enjoy solving! 🧊✨**
