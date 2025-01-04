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
	router.POST("/is_correct", engine.IsGridCorrectAPI)
	router.POST("/getsolution", engine.OneSolutionAPI)
	router.POST("/solve", engine.SolveAPI)
	router.POST("/getallsolutions", engine.AllMoovesAPI)
	router.POST("/all_next_mooves", engine.AllNextMoovesAPI)
	router.POST("/check_grid",engine.CheckCustom)
	
	router.Run("localhost:8081")

}
