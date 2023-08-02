package main

import (
	"fmt"
	"log"
	"os"

	"player_backend/controllers"
	"player_backend/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.Migrate()
}

func main(){
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		fmt.Fprintf(c.Writer, "<h1 style='text-align:center;'>This is player's API</h1>")
	})
	router.POST("/players", controllers.PlayerCreate)
	router.PUT("/players/:id", controllers.UpdatePlayer)
	router.DELETE("/players/:id", controllers.DeletePlayer)
	router.GET("/players", controllers.ListPlayers)
	router.GET("/players/rank/:val", controllers.RankPlayers)
	router.GET("/players/random", controllers.RandomPlayer)

	if err := router.Run(":" + os.Getenv("PORT")); err != nil{
		log.Fatal(err)
	}
}