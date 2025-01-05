package main

import (
	"back/api/engine"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))
	router.POST("/api/solver/is_correct", engine.IsGridCorrectAPI)
	router.POST("/api/solver/getsolution", engine.OneSolutionAPI)
	router.POST("/api/solver/solve", engine.SolveAPI)
	router.POST("/api/solver/getallsolutions", engine.AllMoovesAPI)
	router.POST("/api/solver/all_next_mooves", engine.AllNextMoovesAPI)
	router.POST("/api/solver/check_grid", engine.CheckCustom)

	router.Run("localhost:8081")

}
