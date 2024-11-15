package engine

import (
	"fmt"
)

type Grid struct {
	Grid [9][9]int8 `json:"grid"`
}

type ResponseSolve struct {
	Success bool       `json:"success"`
	Grid    [9][9]int8 `json:"grid"`
	Message string     `json:"message"`
}

type ResponseSolution struct {
	Solution Solution `json:"solution"`
	Grid     Grid     `json:"grid"`
}

type Solution struct {
	Value  int8   `json:"value"`
	Row    int8   `json:"row"`
	Col    int8   `json:"col"`
	Method string `json:"method"`
	Type   string `json:"type"`
}

func DisplayGrid(grid *Grid) {
	for _, row := range grid.Grid {
		for _, col := range row {
			print(col)
		}
		print("\n")
	}
	print("\n")
}

func IsNumberInRow(grid *Grid, row int8, num int8) bool {
	for _, val := range grid.Grid[row] {
		if val == num {
			return true
		}
	}
	return false
}

func IsNumberInCol(grid *Grid, col int8, num int8) bool {
	for i := range grid.Grid {
		if grid.Grid[i][col] == num {
			return true
		}
	}
	return false
}

func IsNumberInBloc(grid *Grid, bloc int8, num int8) bool {

	startRow := 3 * (bloc / 3)
	startCol := 3 * (bloc % 3)

	for i := startRow; i <= startRow+2; i++ {
		for j := startCol; j <= startCol+2; j++ {
			if grid.Grid[i][j] == num {
				return true
			}
		}
	}
	return false
}

func IsGridValueCorrect(grid *Grid) ResponseSolve {
	print(len(grid.Grid))
	for _, row := range grid.Grid {
		for _, val := range row {

			if val < 0 || val > 9 {
				return ResponseSolve{false, grid.Grid, fmt.Sprintf("Invalid value in Grid , found %d", val)}
			}
		}
	}

	return ResponseSolve{true, grid.Grid, ""}
}

func isGridSolvable(grid *Grid) bool {

	return true
}

func isGridValid(grid *Grid) {

}
func BruteForceSolveCopy(result *Grid) ([9][9]int8, bool) {
	var row, col int8
	for row = 0; row < 9; row++ {
		for col = 0; col < 9; col++ {
			if result.Grid[row][col] == 0 {
				for num := int8(1); num <= 9; num++ {
					if IsValid(result, row, col, num) {
						result.Grid[row][col] = num
						if _, ok := BruteForceSolve(result); ok {
							return result.Grid, true
						}
						result.Grid[row][col] = 0
					}

				}
				return result.Grid, false
			}
		}
	}

	return result.Grid, true

}
func BruteForceSolve(grid *Grid) ([9][9]int8, bool) {
	result := grid
	if _, ok := BruteForceSolveCopy(result); ok {
		return result.Grid, true
	}
	return grid.Grid, false

}

func IsValid(grid *Grid, row int8, col int8, num int8) bool {
	for i := 0; i < 9; i++ {
		if grid.Grid[row][i] == num {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if grid.Grid[i][col] == num {
			return false
		}
	}

	startRow, startCol := (row/3)*3, (col/3)*3
	var i, j int8
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			if grid.Grid[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}

func FindBlocFromCoordinate(row int8, col int8) int8 {
	return 3*(row/3) + (col / 3)
}
func CreateCopy(grid *Grid) Grid {
	copy := Grid{}
	var i, j int8
	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			copy.Grid[i][j] = grid.Grid[i][j]
		}
	}
	return copy
}
