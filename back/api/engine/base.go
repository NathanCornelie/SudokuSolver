package engine

import (
	"fmt"
)

// FindPlacesInBloc Finds all empty positions in a specific 3x3 block.
func FindPlacesInBloc(grid *Grid, bloc int8) []struct{ row, col int8 } {
	rowStart := 3 * (bloc / 3)
	colStart := 3 * (bloc % 3)
	var FindPositions []struct{ row, col int8 }

	for j := rowStart; j < rowStart+3; j++ {
		for k := colStart; k < colStart+3; k++ {
			if grid.Grid[j][k] == 0 {
				FindPositions = append(FindPositions, struct{ row, col int8 }{j, k})
			}
		}
	}
	return FindPositions
}

// FindPlacesInRow Finds all empty positions in a specific row.
func FindPlacesInRow(grid *Grid, row int8) []struct{ row, col int8 } {
	var FindPositions []struct{ row, col int8 }
	var ind int8
	for ind = 0; ind < 9; ind++ {
		if grid.Grid[row][ind] == 0 {
			FindPositions = append(FindPositions, struct{ row, col int8 }{row, ind})
		}
	}
	return FindPositions
}

// FindPlacesInCol Finds all empty positions in a specific column.
func FindPlacesInCol(grid *Grid, col int8) []struct{ row, col int8 } {
	var FindPositions []struct{ row, col int8 }
	var ind int8
	for ind = 0; ind < 9; ind++ {
		if grid.Grid[ind][col] == 0 {
			FindPositions = append(FindPositions, struct{ row, col int8 }{ind, col})
		}
	}
	return FindPositions
}

func FillEmptyPlacesInBloc(grid *Grid, bloc int8) {

	rowStart := 3 * (bloc / 3)
	colStart := 3 * (bloc % 3)
	for row := rowStart; row < rowStart+2; row++ {
		for col := colStart; col < colStart+2; col++ {
			if grid.Grid[row][col] == 0 {
				grid.Grid[row][col] = 10
			}
		}
	}
}
func FillEmptyPlacesInCol(grid *Grid, col int8) {
	var row int8
	for row = 0; row < 9; row++ {
		if grid.Grid[row][col] == 0 {
			grid.Grid[row][col] = 10

		}
	}
}
func FillEmptyPlacesInRow(grid *Grid, row int8) {
	var col int8
	for col = 0; col < 9; col++ {
		if grid.Grid[row][col] == 0 {
			grid.Grid[row][col] = 10

		}
	}
}

// FindSimpleNumber Finds and places a given number if it has a unique spot in a block, row, or column.
func FindSimpleNumber(grid *Grid, cible int8, manquants *[9]int8, numberPlaced *bool, solutions []Solution, one bool) []Solution {
	var copy = CreateCopy(grid)

	// Mark cells where the target number cannot go
	for col := 0; col < 9; col++ {
		for row := 0; row < 9; row++ {
			if copy.Grid[row][col] == cible {
				for ind := 0; ind < 9; ind++ {
					if copy.Grid[row][ind] == 0 {
						copy.Grid[row][ind] = 10 // Mark the cell as unavailable
					}
					if copy.Grid[ind][col] == 0 {
						copy.Grid[ind][col] = 10 // Mark the cell as unavailable
					}
				}
			}
		}
	}

	var i, j int8
	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			positionsDisponibles := FindPlacesInBloc(&copy, j)
			if len(positionsDisponibles) == 1 {
				row, col := positionsDisponibles[0].row, positionsDisponibles[0].col
				if !IsNumberInBloc(&copy, j, cible) {
					copy.Grid[row][col] = cible
					FillEmptyPlacesInBloc(&copy, j)
					FillEmptyPlacesInCol(&copy, j)
					FillEmptyPlacesInRow(&copy, j)

				}
				if !IsNumberInBloc(grid, j, cible) {
					grid.Grid[row][col] = cible
					*numberPlaced = true
					manquants[cible-1]--

					solutions = append(solutions, Solution{cible, row, col, "base", "bloc"})
					fmt.Printf("%d placed in position %d,%d (bloc)\n", cible, row, col)
					if one {
						return solutions
					}
				}
			}

			positionsDisponibles = FindPlacesInRow(&copy, i)
			if len(positionsDisponibles) == 1 {
				row, col := positionsDisponibles[0].row, positionsDisponibles[0].col
				if !IsNumberInRow(&copy, i, cible) {
					copy.Grid[row][col] = cible
					FillEmptyPlacesInBloc(&copy, j)
					FillEmptyPlacesInCol(&copy, j)
					FillEmptyPlacesInRow(&copy, j)
				}
				if !IsNumberInRow(grid, i, cible) {

					grid.Grid[row][col] = cible
					*numberPlaced = true
					manquants[cible-1]--
					solutions = append(solutions, Solution{cible, row, col, "base", "row"})
					fmt.Printf("%d placed in position %d,%d (row)\n", cible, row, col)
					if one {
						return solutions
					}
				}
			}

			positionsDisponibles = FindPlacesInCol(&copy, i)
			if len(positionsDisponibles) == 1 {
				row, col := positionsDisponibles[0].row, positionsDisponibles[0].col
				if !IsNumberInCol(&copy, i, cible) {
					copy.Grid[row][col] = cible
					FillEmptyPlacesInBloc(&copy, j)
					FillEmptyPlacesInCol(&copy, j)
					FillEmptyPlacesInRow(&copy, j)
				}
				if !IsNumberInCol(grid, i, cible) {
					grid.Grid[row][col] = cible
					*numberPlaced = true
					manquants[cible-1]--
					solutions = append(solutions, Solution{cible, row, col, "base", "col"})
					fmt.Printf("%d placed in position %d,%d (col)\n", cible, row, col)
					if one {
						return solutions
					}
				}
			}
		}
	}
	return solutions
}
