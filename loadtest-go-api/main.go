package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"loadtest-go-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")

    dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=Asia/Bangkok"
    var err error
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    db.AutoMigrate(&models.Item{})
}

func createItem(c *gin.Context) {
    var newItem models.Item
    if err := c.ShouldBindJSON(&newItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.Create(&newItem)
    c.JSON(http.StatusCreated, newItem)
}

func getItem(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    var item models.Item
    if result := db.First(&item, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
        return
    }
    c.JSON(http.StatusOK, item)
}

func getItems(c *gin.Context) {
    var items []models.Item
    db.Find(&items)
    c.JSON(http.StatusOK, items)
}

func updateItem(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var updatedItem models.Item
    if err := c.ShouldBindJSON(&updatedItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var item models.Item
    if result := db.First(&item, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
        return
    }

    item.Name = updatedItem.Name
    db.Save(&item)
    c.JSON(http.StatusOK, item)
}

func deleteItem(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if result := db.Delete(&models.Item{}, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
        return
    }
    c.JSON(http.StatusNoContent, gin.H{})
}

func main() {
    initDB()

    r := gin.Default()

    r.POST("/items", createItem)
    r.GET("/items/:id", getItem)
    r.GET("/items", getItems)
    r.PUT("/items/:id", updateItem)
    r.DELETE("/items/:id", deleteItem)

    r.Run(":8080")
}
