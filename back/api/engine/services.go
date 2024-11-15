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
	var solutions []Solution
	solutions = Solver(&grid, solutions, true)

	if len(solutions) > 0 {

		c.JSON(http.StatusOK, ResponseSolution{solutions[0], grid})
	} else {
		c.JSON(http.StatusInternalServerError, "no solution founded")
	}

}
func AllMoovesAPI(c *gin.Context) {
	return
}
func AllNextMoovesAPI(c *gin.Context) {
	return
}
