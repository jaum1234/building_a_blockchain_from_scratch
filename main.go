package main

import (
	"encoding/hex"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	bc := NewBlockchain()

	engine := gin.Default()

	engine.GET("/blocks", func(c *gin.Context) {
		blocks := make([]map[string]any, 0)

		for _, v := range bc.blocks {
			blocks = append(blocks, map[string]any{
				"Timestamp":        v.Timestamp,
				"Data":             string(v.Data),
				"Hash":             hex.EncodeToString(v.Hash),
				"Previous Hash":    hex.EncodeToString(v.PrevBlockHash),
				"Number used Once": v.Nonce,
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
				"Timestamp":        b.Timestamp,
				"Data":             string(b.Data),
				"Hash":             hex.EncodeToString(b.Hash),
				"Previous Hash":    hex.EncodeToString(b.PrevBlockHash),
				"Number used Once": b.Nonce,
			},
		})
	})

	engine.Run()
}
