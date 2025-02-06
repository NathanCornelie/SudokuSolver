package engine

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func IsGridCorrectAPI(c *gin.Context) {
	var grid Grid
	if err := c.BindJSON(&grid); err != nil {
		print(err)
		c.JSON(http.StatusBadRequest, ResponseSolve{false, grid.Grid, "Invalid grid"})
		return
	}

	resp := IsGridValueCorrect(&grid)
	if !resp.Success {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	result, ok := BruteForceSolve(&grid)
	if ok {
		c.JSON(http.StatusBadRequest, ResponseSolve{true, result, "Valid grid sovled!!"})
		return

	} else {
		c.JSON(http.StatusOK, ResponseSolve{false, result, "There is no valid answer for this grid"})
		return
	}

}

func SolveAPI(c *gin.Context) {
	var grid Grid
	if err := c.BindJSON(&grid); err != nil {
		print(err)
		c.JSON(http.StatusBadRequest, ResponseSolve{false, grid.Grid, "Invalid grid"})
		return
	}

	resp := IsGridValueCorrect(&grid)
	if !resp.Success {
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	copy := CreateCopy(&grid)

	result, ok := BruteForceSolve(&copy)
	if !ok {
		c.JSON(http.StatusBadRequest, ResponseSolve{false, result, "There is no valid answer for this grid"})
		return
	} else {

		c.JSON(http.StatusOK, ResponseSolve{true, result, ""})
	}
	// solutions := Solver(&grid)
	// c.JSON(http.StatusOK, solutions)

}

func OneSolutionAPI(c *gin.Context) {
	var grid Grid
	if err := c.BindJSON(&grid); err != nil {
		print(err)
		c.JSON(http.StatusBadRequest, ResponseSolve{false, grid.Grid, "Invalid grid"})
		return
	}

	resp := IsGridValueCorrect(&grid)
	if !resp.Success {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	solver := Solver(&grid, true)
	solutions := solver.Solutions
	if len(solutions) > 0 {

		c.JSON(http.StatusOK, solutions[0])
	} else {
		c.JSON(http.StatusOK, solutions)
	}

}
func AllMoovesAPI(c *gin.Context) {
	var grid Grid
	if err := c.BindJSON(&grid); err != nil {
		print(err)
		c.JSON(http.StatusBadRequest, ResponseSolve{false, grid.Grid, "Invalid grid"})
		return
	}

	resp := IsGridValueCorrect(&grid)
	if !resp.Success {
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	solver := Solver(&grid, false)
	solutions := solver.Solutions
	if len(solutions) > 0 {

		c.JSON(http.StatusOK, solutions)
	} else {
		c.JSON(http.StatusOK, solutions)
	}
}


func CheckCustom(c *gin.Context) {
	var grid Grid
	if err := c.BindJSON(&grid); err != nil {
		print(err)
		c.JSON(http.StatusBadRequest, ResponseSolve{false, grid.Grid, "Invalid grid"})
		return
	}

	resp := IsGridValueCorrect(&grid)
	if !resp.Success {
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	valid := isGridValid(&grid)
	if !valid {
		c.JSON(http.StatusBadRequest, ResponseSolve{false, grid.Grid, "Invalid grid"})
		return
	}
	solvable := isGridSolvable(&grid)
	if !solvable {
		c.JSON(http.StatusBadRequest, ResponseSolve{false, grid.Grid, "Invalid grid"})
	}

	solver := Solver(&grid, false)

	if solver.Completed && len(solver.Solutions) > 0 {

		c.JSON(http.StatusOK, ResponseSolve{true, grid.Grid, "Solutions have been found"})

	} else {
		c.JSON(http.StatusBadRequest, ResponseSolve{false, grid.Grid, "We did not manage to solve that grid :("})
	}

}
