#!/usr/bin/env python3
"""
Kociemba Rubik's Cube Solver Wrapper
Accepts cube state as 54-character string, returns optimal solution
"""

import sys
import kociemba

def solve_cube(cube_string):
    """
    Solve a Rubik's cube using Kociemba's algorithm

    Args:
        cube_string: 54-character string representing cube state
                    Format: UUUUUUUUURRRRRRRRRFFFFFFFFFDDDDDDDDDLLLLLLLLLBBBBBBBBB
                    Order: Up(9), Right(9), Front(9), Down(9), Left(9), Back(9)

    Returns:
        Solution string with space-separated moves (e.g., "R U R' U'")
    """
    try:
        solution = kociemba.solve(cube_string)
        return solution
    except Exception as e:
        return f"ERROR: {str(e)}"

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: kociemba_solver.py <54-char-cube-string>")
        sys.exit(1)

    cube_string = sys.argv[1]

    if len(cube_string) != 54:
        print(f"ERROR: Cube string must be exactly 54 characters, got {len(cube_string)}")
        sys.exit(1)

    solution = solve_cube(cube_string)
    print(solution)
