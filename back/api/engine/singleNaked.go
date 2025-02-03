package engine

import (
	"fmt"
)

// findSingleNaked identifies and places numbers when only one possibility exists in a row.
func FindSingleNaked(grid *Grid, numberPlaced *bool) []ResponseSolution {
	var row int8
	var solutions []ResponseSolution
	for row = 0; row < 9; row++ {
		manquants := make(map[int8]struct{}, 9)
		presents := make(map[int8]struct{}, 9)

		// Collect present numbers in the row
		for col := 0; col < 9; col++ {
			if grid.Grid[row][col].Value != 0 {
				presents[grid.Grid[row][col].Value] = struct{}{}
			}
		}

		var col int8
		for col = 1; col <= 9; col++ {
			if _, exists := presents[col]; !exists {
				manquants[col] = struct{}{}
			}
		}

		for col = 0; col < 9; col++ {
			if grid.Grid[row][col].Value == 0 {
				indiceBloc := FindBlocFromCoordinate(row, col)
				possibles := make(map[int8]struct{})
				if len(manquants) == 1 {
					for num := range manquants {
						grid.Grid[row][col].Value = num
						*numberPlaced = true
						fmt.Println("Single Naked")

						solutions = append(solutions, ResponseSolution{Solution{num, row, col, "Single Naked", "row"}, CreateCopy(grid)})

					}
				}
				// Copy manquants into possibles
				var num int8
				for num = range manquants {
					possibles[num] = struct{}{}
				}

				// Remove numbers that are already in the block or column
				for num := range possibles {

					if IsNumberInCol(grid, col, num) {
						delete(possibles, num)
					}
				}
				if len(possibles) == 1 {
					for num := range possibles {
						grid.Grid[row][col].Value = num
						*numberPlaced = true
						fmt.Println("Single Naked")

						solutions = append(solutions, ResponseSolution{Solution{num, row, col, "Single Naked", "col"}, CreateCopy(grid)})

					}
				} else {
					for num := range possibles {

						if IsNumberInBloc(grid, indiceBloc, num) || IsNumberInCol(grid, col, num) {
							delete(possibles, num)
						}
					}

					if len(possibles) == 1 {
						for num := range possibles {
							grid.Grid[row][col].Value = num
							*numberPlaced = true
							fmt.Println("Single Naked")

							solutions = append(solutions, ResponseSolution{Solution{num, row, col, "Single Naked", ""}, CreateCopy(grid)})

						}
					}
				}

			}
		}
	}
	return solutions
}

// Helper functions (isNumberInBlc, isNumberInCol, findBlocFromCoordinate, and placeNumberInGrid) need to be implemented.
