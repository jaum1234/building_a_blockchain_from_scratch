package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	bc := NewBlockchain()

	engine := gin.Default()

	engine.GET("/blocks", func(c *gin.Context) {
		blocks := make([]map[string]any, 0)

		for _, v := range bc.blocks {
			blocks = append(blocks, map[string]any{
				"Timestamp":     v.Timestamp,
				"Data":          string(v.Data),
				"Hash":          v.Hash,
				"Previous Hash": v.PrevBlockHash,
			})
		}

		c.JSON(200, gin.H{
			"success": true,
			"message": "Blocks listed",
			"data":    blocks,
		})
	})

	engine.POST("/blocks", func(c *gin.Context) {
		var block BlockToBeAdded

		if err := c.BindJSON(&block); err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		bc.AddBlock(block.Data)
		b := bc.LastBlock()

		c.JSON(200, gin.H{
			"success": true,
			"message": "Block inserted",
			"data": map[string]any{
				"Timestamp":     b.Timestamp,
				"Data":          string(b.Data),
				"Hash":          b.Hash,
				"Previous Hash": b.PrevBlockHash,
			},
		})
	})

	engine.Run(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT"))
}
