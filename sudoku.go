// sudoku is a simple sudoku puzzle solver.
// It uses a simple (brute force) generate and test search, with the cell selection being constrained
// by the current content of the relevant row, column, and local square.
package main

import "fmt"

type suGrid [9][9]uint8

var grid suGrid

func solve() bool {
	var v uint8
	for ri, row := range grid {
		for ci, _ := range row {
			if grid[ri][ci] != 0 {
				continue
			}
			pvs := posblVals(ri, ci)
			if len(pvs) == 0 {
				return false
			}
			for _, v = range pvs {
				grid[ri][ci] = v
				solved := solve()
				if solved {
					return true
				}
			}
			grid[ri][ci] = 0
			return false
		}
	}
	fmt.Println("Solution")
	grid.print()
	return true
}

func (g suGrid) print() {
	for i, c := range g {
		for j := range c {
			fmt.Printf("%v ", g[i][j])
		}
		fmt.Printf("\n")
	}
}

// posblVals finds the digits that can be inserted at i,j by firstly setting up a map of the digits in row i,
// column j and the "local square" (one of the 9 squares in the grid), then finding the digits not in the map.
// There is a very small amount of slightly untidy duplication in this processing.
// The result is a (possibly empty) slice of digits.
func posblVals(row int, col int) []uint8 {
	var vals []uint8
	used := make(map[uint8]uint8)
	var c uint8
	for _, c = range grid[row] {
		used[c] = 1
	}
	for j := 0; j < 9; j++ {
		used[grid[j][col]] = 1
	}
	// search the square containg i,j.
	for i := row / 3 * 3; i < row/3*3+3; i++ {
		for j := col / 3 * 3; j < col/3*3+3; j++ {
			used[grid[i][j]] = 1
		}
	}
	for c = 1; c <= 9; c++ {
		if _, ok := used[c]; !ok {
			vals = append(vals, c)
		}
	}
	return vals
}

func main() {
	grid.setup()
	fmt.Println("Problem")
	grid.print()
	solve()
}

// empty cells are left set at zero.
func (g *suGrid) setup() {
	g[0][4] = 4
	g[0][6] = 9
	g[1][2] = 9
	g[1][4] = 6
	g[1][7] = 2
	g[1][8] = 1
	g[2][1] = 4
	g[2][5] = 7
	g[2][8] = 6
	g[3][0] = 9
	g[3][2] = 3
	g[3][4] = 2
	g[4][1] = 8
	g[4][3] = 4
	g[4][5] = 9
	g[4][7] = 1
	g[5][4] = 8
	g[5][6] = 6
	g[5][8] = 2
	g[6][0] = 1
	g[6][3] = 8
	g[6][7] = 5
	g[7][0] = 6
	g[7][1] = 3
	g[7][4] = 5
	g[7][6] = 1
	g[8][2] = 7
	g[8][4] = 1
}

// Another test case.
//func (g *suGrid) setup() {
//	g[0][0] = 4
//	g[0][4] = 6
//	g[0][5] = 8
//	g[0][8] = 5
//	g[1][3] = 7
//	g[1][4] = 3
//	g[1][5] = 9
//	g[1][8] = 8
//	g[2][1] = 6
//	g[2][7] = 1
//	g[2][8] = 7
//	g[3][0] = 3
//	g[3][7] = 7
//	g[3][8] = 1
//	g[4][2] = 2
//	g[4][4] = 1
//	g[4][6] = 4
//	g[5][0] = 5
//	g[5][1] = 7
//	g[5][8] = 2
//	g[6][0] = 1
//	g[6][1] = 2
//	g[6][7] = 9
//	g[7][0] = 7
//	g[7][3] = 9
//	g[7][4] = 2
//	g[7][5] = 3
//	g[8][0] = 8
//	g[8][3] = 6
//	g[8][4] = 7
//	g[8][8] = 4
//}
