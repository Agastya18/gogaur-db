package main

import (
	"net/http"
	"os"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"saaster.tech/own-db/db"
)

var database *db.SimpleDB

func main() {
	// Initialize the database
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading env")
	}
	var err error
	database, err = db.OpenDB("mydb.data")
	
	if err != nil {
		panic("Failed to open database: " + err.Error())
	}
	defer database.Close()
    PORT := os.Getenv("PORT")
	if PORT == "" {
        PORT = "8080" // Default port
    }
	r := gin.Default()

	r.POST("/set", handleSet)
	r.GET("/get", handleGet)
	r.DELETE("/delete", handleDelete)

	if err := r.Run(":" + PORT); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

func handleSet(c *gin.Context) {
	var body struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := database.Set(body.Key, body.Value); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Key set successfully"})
}

func handleGet(c *gin.Context) {
	key := c.Query("key")
	value, err := database.Get(key)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"key": key, "value": value})
}

func handleDelete(c *gin.Context) {
	key := c.Query("key")
	err := database.Delete(key)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
		return
	}

	c.Status(http.StatusOK)
}
