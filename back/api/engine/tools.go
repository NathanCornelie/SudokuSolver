package engine

import (
	"fmt"
)

type Case struct {
	Value int8   `json:"value"`
	Color string `json:"color"`
}

type Grid struct {
	Grid [9][9]Case `json:"grid"`
}

type ResponseSolve struct {
	Success bool       `json:"success"`
	Grid    [9][9]Case `json:"grid"`
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
			print(col.Value)
		}
		print("\n")
	}
	print("\n")
}

func IsNumberInRow(grid *Grid, row int8, num int8) bool {
	for _, val := range grid.Grid[row] {
		if val.Value == num {
			return true
		}
	}
	return false
}

func IsNumberInCol(grid *Grid, col int8, num int8) bool {
	for i := range grid.Grid {
		if grid.Grid[i][col].Value == num {
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
			if grid.Grid[i][j].Value == num {
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

			if val.Value < 0 || val.Value > 9 {
				return ResponseSolve{false, grid.Grid, fmt.Sprintf("Invalid value in Grid , found %d", val.Value)}
			}
		}
	}

	return ResponseSolve{true, grid.Grid, ""}
}

func isGridSolvable(grid *Grid) bool {
	fmt.Println("bruteforce")
	_, resp := BruteForceSolve(grid)

	return resp
}

func BruteForceSolveCopy(result *Grid) ([9][9]Case, bool) {
	var row, col int8
	for row = 0; row < 9; row++ {
		for col = 0; col < 9; col++ {
			if result.Grid[row][col].Value == 0 {
				for num := int8(1); num <= 9; num++ {
					if IsValid(result, row, col, num) {
						result.Grid[row][col].Value = num
						if _, ok := BruteForceSolve(result); ok {
							return result.Grid, true
						}
						result.Grid[row][col].Value = 0
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
		column := [9]Case{}
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
			block := [9]Case{}
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
func isValidSet(nums []Case) bool {
	seen := make(map[int8]bool)
	for _, num := range nums {
		if num.Value != 0 { // Ignorer les cases vides
			if seen[num.Value] {
				return false
			}
			seen[num.Value] = true
		}
	}
	return true
}
func BruteForceSolve(grid *Grid) ([9][9]Case, bool) {

	result := grid
	if _, ok := BruteForceSolveCopy(result); ok {
		return result.Grid, true
	}
	return grid.Grid, false

}

func IsValid(grid *Grid, row int8, col int8, num int8) bool {
	for i := 0; i < 9; i++ {
		if grid.Grid[row][i].Value == num {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if grid.Grid[i][col].Value == num {
			return false
		}
	}

	startRow, startCol := (row/3)*3, (col/3)*3
	var i, j int8
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			if grid.Grid[startRow+i][startCol+j].Value == num {
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

func ClearColor(grid *Grid) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			grid.Grid[i][j].Color = ""
		}
	}
}

func ColorBloc(grid *Grid, bloc int8, value int8) {
	startRow := 3 * (bloc / 3)
	startCol := 3 * (bloc % 3)

	for i := startRow; i <= startRow+2; i++ {
		for j := startCol; j <= startCol+2; j++ {
			if grid.Grid[i][j].Value == value {
				grid.Grid[i][j].Color = "red1"
			} else {
				grid.Grid[i][j].Color = "red2"
			}
		}
	}
}
func ColorRow(grid *Grid, row int8, value int8) {
	for i := 0; i < 9; i++ {
		if grid.Grid[row][i].Value == value {
			grid.Grid[row][i].Color = "red1"
		} else {
			grid.Grid[row][i].Color = "red2"
		}
	}
}
func ColorCol(grid *Grid, col int8, value int8) {
	for i := 0; i < 9; i++ {
		if grid.Grid[i][col].Value == value {
			grid.Grid[i][col].Color = "red1"
		} else {
			grid.Grid[i][col].Color = "red2"
		}
	}
}
func ColorSimpleCol(grid *Grid, pos struct{ row, col int8 }) {
	var value = grid.Grid[pos.row][pos.col].Value
	var row int8
	ClearColor(grid)
	for row = 0; row < 9; row++ {
		if grid.Grid[row][pos.col].Value == 0 {
			var bloc = FindBlocFromCoordinate(row, pos.col)
			if IsNumberInBloc(grid, bloc, value) && bloc != FindBlocFromCoordinate(pos.row, pos.col) {
				ColorBloc(grid, bloc, value)
			} else {
				ColorRow(grid, row, value)
			}
		}
	}

}
func ColorSimpleRow(grid *Grid, pos struct{ row, col int8 }) {
	var value = grid.Grid[pos.row][pos.col].Value
	var col int8
	ClearColor(grid)
	for col = 0; col < 9; col++ {
		if grid.Grid[pos.row][col].Value == 0 {
			var bloc = FindBlocFromCoordinate(pos.row, col)
			if IsNumberInBloc(grid, bloc, value) && bloc != FindBlocFromCoordinate(pos.row, pos.col) {
				ColorBloc(grid, bloc, value)
			} else {
				ColorCol(grid, col, value)
			}
		}
	}
}

func ColorSimpleBloc(grid *Grid, pos struct{ row, col int8 }) {
	var value = grid.Grid[pos.row][pos.col].Value
	var bloc = FindBlocFromCoordinate(pos.row, pos.col)
	startRow := 3 * (bloc / 3)
	startCol := 3 * (bloc % 3)
	ClearColor(grid)
	for i := startRow; i <= startRow+2; i++ {
		for j := startCol; j <= startCol+2; j++ {
			if grid.Grid[i][j].Value == 0 {
				if IsNumberInCol(grid, j, value) && j != pos.col {
					fmt.Printf("YES col %d \n", j)
					ColorCol(grid, j, value)
				}
				if IsNumberInRow(grid, i, value) && i != pos.row {
					fmt.Printf("YES row %d \n", i)
					ColorRow(grid, i, value)
				}
			}

		}
	}
}
func ColorSingleNakedRow(grid *Grid, row int8, val int8) {
	ClearColor(grid)
	for k := 0; k < 9; k++ {
		if grid.Grid[row][k].Value != 0 {
			grid.Grid[row][k].Color = "red1"
		}
	}
}

func ColorSingleNakedCol(grid *Grid, pos struct{ row, col int8 }, toColor []int8) {
	ClearColor(grid)
	for k := 0; k < 9; k++ {
		if grid.Grid[pos.row][k].Value != 0 {
			grid.Grid[pos.row][k].Color = "red1"
		}

		for _, v := range toColor {
			if v == grid.Grid[k][pos.col].Value {
				grid.Grid[k][pos.col].Color = "red1"
			}
		}
		grid.Grid[pos.row][pos.col].Color = "red1"

	}
}
func ColorSingleNakedGlobal(grid *Grid, pos struct{ row, col int8 }, toColorBloc []int8, toColorCol []int8) {
	ClearColor(grid)

	for k := 0; k < 9; k++ {
		if grid.Grid[pos.row][k].Value != 0 {
			grid.Grid[pos.row][k].Color = "red1"
		}

		for _, v := range toColorCol {
			fmt.Print(toColorCol, grid.Grid[k][pos.col].Value, toColorBloc, v, pos, "\n")
			if v == grid.Grid[k][pos.col].Value {

				grid.Grid[k][pos.col].Color = "red1"
			}
		}
	}
	bloc := FindBlocFromCoordinate(pos.row, pos.col)
	startRow := 3 * (bloc / 3)
	startCol := 3 * (bloc % 3)

	for i := startRow; i <= startRow+2; i++ {
		for j := startCol; j <= startCol+2; j++ {
			for _, v := range toColorBloc {
				if v == grid.Grid[i][j].Value {
					grid.Grid[i][j].Color = "red1"
				}
			}
		}

	}
}
