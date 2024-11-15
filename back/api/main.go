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
	router.POST("/getsolutions", engine.AllMoovesAPI)
	router.POST("/all_next_moooves", engine.AllNextMoovesAPI)

	router.Run("localhost:8080")

}
