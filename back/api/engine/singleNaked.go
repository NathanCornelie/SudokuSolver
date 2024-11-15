package engine

import (
	"fmt"
)

// findSingleNaked identifies and places numbers when only one possibility exists in a row.
func FindSingleNaked(grid *Grid, numberPlaced *bool, solutions []Solution, one bool) []Solution {
	var row int8
	for row = 0; row < 9; row++ {
		manquants := make(map[int8]struct{}, 9)
		presents := make(map[int8]struct{}, 9)

		// Collect present numbers in the row
		for col := 0; col < 9; col++ {
			if grid.Grid[row][col] != 0 {
				presents[grid.Grid[row][col]] = struct{}{}
			}
		}

		var col int8
		for col = 1; col <= 9; col++ {
			if _, exists := presents[col]; !exists {
				manquants[col] = struct{}{}
			}
		}

		for col = 0; col < 9; col++ {
			if grid.Grid[row][col] == 0 {
				indiceBloc := FindBlocFromCoordinate(row, col)
				possibles := make(map[int8]struct{})

				// Copy manquants into possibles
				var num int8
				for num = range manquants {
					possibles[num] = struct{}{}
				}

				// Remove numbers that are already in the block or column
				for num := range possibles {
					if IsNumberInBloc(grid, indiceBloc, num) || IsNumberInCol(grid, col, num) {
						delete(possibles, num)
					}
				}

				if len(possibles) == 1 {
					for num := range possibles {
						grid.Grid[row][col] = num
						*numberPlaced = true
						if one {
							solutions = append(solutions, Solution{num, row, col, "singleNaked", ""})
							return solutions
						}
						fmt.Println("Single Naked")
					}
				}
			}
		}
	}
	return solutions
}

// Helper functions (isNumberInBlc, isNumberInCol, findBlocFromCoordinate, and placeNumberInGrid) need to be implemented.
