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

type SolverSolution struct {
	Completed bool               `json:"completed"`
	Solutions []ResponseSolution `json:"solution"`
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
	fmt.Println("check values")
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
	_,resp := BruteForceSolve(grid);

	return resp
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

func isGridValid(grid *Grid) bool {
	fmt.Println("check grid")

	// Vérifier les lignes
	for i := 0; i < 9; i++ {
		if !isValidSet(grid.Grid[i][:]) {
			return false
		}
	}

	// Vérifier les colonnes
	for i := 0; i < 9; i++ {
		column := [9]int8{}
		for j := 0; j < 9; j++ {
			column[j] = grid.Grid[j][i]
		}
		if !isValidSet(column[:]) {
			return false
		}
	}

	// Vérifier les blocs 3x3
	for row := 0; row < 9; row += 3 {
		for col := 0; col < 9; col += 3 {
			block := [9]int8{}
			idx := 0
			for i := row; i < row+3; i++ {
				for j := col; j < col+3; j++ {
					block[idx] = grid.Grid[i][j]
					idx++
				}
			}
			if !isValidSet(block[:]) {
				return false
			}
		}
	}

	return true
}
func isValidSet(nums []int8) bool {
	seen := make(map[int8]bool)
	for _, num := range nums {
		if num != 0 { // Ignorer les cases vides
			if seen[num] {
				return false
			}
			seen[num] = true
		}
	}
	return true
}
func BruteForceSolve(grid *Grid) ([9][9]int8, bool) {
	fmt.Println("bruteforce")
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
