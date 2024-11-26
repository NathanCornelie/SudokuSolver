package engine

import (
	"fmt"
)

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
func IsListEmpty(list *[9]int8) bool {
	fmt.Printf("\nnathan\n")
	var k int8
	for k = 0; k < 9; k++ {
		fmt.Printf("%d", list[k])
		if list[k] != 0 {
			return false
		}
	}
	return true
}

func Solver(grid *Grid, one bool) []ResponseSolution {
	nbrChiffreManquant := [9]int8{}
	oldNbrChiffreManquant := [9]int8{}
	InitManquants(grid, &nbrChiffreManquant)
	solutions := []ResponseSolution{}
	var solution ResponseSolution
	numberPlaced := false
	var i int8

	for !IsListEmpty(&nbrChiffreManquant) {
		for j := 0; j < 9; j++ {
			for i = 0; i < 9; i++ {
				numberPlaced = false
				if nbrChiffreManquant[i] > 0 {
					solutions = append(solutions, FindSimpleNumber(grid, i+1, &nbrChiffreManquant, &numberPlaced)...)
					if one && numberPlaced && len(solutions) > 0 {
						solutions = append(solutions, solution)
						return solutions
					}
					// if !numberPlaced {
					// 	solutions = append(solutions, FindSingleNaked(grid, &numberPlaced)...)
					// 	if one && numberPlaced && len(solutions) > 0 {
					// 		solutions = append(solutions, solution)
					// 		return solutions
					// 	}
					// }
				}

			}
		}
		var k int8
		for k = 0; k < 9; k++ {
			fmt.Printf("%d", nbrChiffreManquant[k])
		}
		if oldNbrChiffreManquant == nbrChiffreManquant {
			return solutions
		}
		oldNbrChiffreManquant = nbrChiffreManquant
	}

	return solutions

}
