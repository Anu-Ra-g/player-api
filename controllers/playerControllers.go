package controllers

import (
	// "fmt"
	"math/rand"
	"net/http"
	"player_backend/initializers"
	"player_backend/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Player struct{
	ID      		uint 					`json:"id" validate:"required"`
	Name    	    string 					`json:"name" validate:"required,max=15"`
	Country 	    string  				`json:"country" validate:"required,uppercase,len=2"`
	Score   		uint 					`json:"score" validate:"required"`
}


func PlayerCreate(c *gin.Context){

	var player Player

	c.Bind(&player)

	validate := validator.New()
	if err := validate.Struct(player); err != nil{
		c.Header("Content-type", "application/json")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad input",
		})
		return
	}

	create_player := models.Player{
		Name: player.Name,
		Country: player.Country,
		Score: player.Score,
	}

	initializers.DB.Create(&create_player)

	c.JSON(200, gin.H{
		"message": create_player,
	})


}
func UpdatePlayer(c *gin.Context){
	var body models.Player

	id := c.Param("id")

	c.Bind(&body)

	var targetPlayer models.Player

	initializers.DB.First(&targetPlayer, id)

	initializers.DB.Model(&targetPlayer).Updates(models.Player{
		Name: body.Name,
		Score: body.Score,
	})

	c.JSON(200, gin.H{
		"post": targetPlayer,
	})


}

func DeletePlayer(c *gin.Context) {
	id := c.Param("id")

	var player models.Player
	if err := initializers.DB.First(&player, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Player not found",
		})
		return
	}

	if err := initializers.DB.Delete(&player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete player",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Player deleted",
	})
}

func ListPlayers(c *gin.Context){

	var players []models.Player
	initializers.DB.Order("name DESC").Find(&players)

	c.JSON(http.StatusOK, gin.H{
		"players": players,
	})
}

func RankPlayers(c *gin.Context) {
	rank, err := strconv.Atoi(c.Param("val"))
	if err != nil || rank <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid rank value",
		})
		return
	}

	var targetPlayer models.Player

	initializers.DB.Order("score DESC").Limit(1).Offset(rank - 1).Find(&targetPlayer)

	c.JSON(http.StatusOK, gin.H{
		"requested_player": targetPlayer,
	})
}


func RandomPlayer(c *gin.Context){
	var random_player models.Player

	count_records := initializers.DB.Find(&[]models.Player{}).RowsAffected

	val := rand.Intn(int(count_records))

	initializers.DB.Order("score").Limit(1).Offset(val - 1).Find(&random_player)

	c.JSON(http.StatusOK, gin.H{
		"random_player": random_player,
	})


}