package engine

func InitManquants(grid *Grid, manquant *[9]int8) {
	finds := [9]int8{0}
	var i int8
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid.Grid[i][j] != 0 {
				finds[grid.Grid[i][j]-1]++
			}
		}
	}
	for i = 0; i < 9; i++ {
		manquant[i] = 9 - finds[i]
	}
}

// solver attempts to solve the Sudoku grid using various strategies.
func Solver(grid *Grid, old_solutions []Solution, one bool) []Solution {
	nbrChiffreManquant := [9]int8{}
	InitManquants(grid, &nbrChiffreManquant)
	solutions := []Solution{}
	numberPlaced := false
	var i int8
	for i = 0; i < 9; i++ {
		if nbrChiffreManquant[i] > 0 {
			solutions = FindSimpleNumber(grid, i+1, &nbrChiffreManquant, &numberPlaced, solutions, one)

			if one && len(solutions) > 0 {
				return solutions
			}
		}
		if !numberPlaced {
			solutions = FindSingleNaked(grid, &numberPlaced, solutions, one)
			if one && len(solutions) > 0 {
				return solutions
			}
		}
	}
	return solutions

}
