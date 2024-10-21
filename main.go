package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album respresents data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albumMap = map[string]album{
	"1": {ID: "1", Title: "Blue Train", Artist: "John Contrane", Price: 56.99},
	"2": {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	"3": {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// get albums reponds with the list of all albums as JSON
func getAlbums(c *gin.Context) {
	var albums []album
	for key, a := range albumMap {
		fmt.Println(key)
		albums = append(albums, a)
	}
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to new Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice
	albumMap[newAlbum.ID] = newAlbum
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumID locates the album whose ID value matches the id
func getAlbumID(c *gin.Context) {
	id := c.Param("id")

	if a, exits := albumMap[id]; exits {
		c.IndentedJSON(http.StatusOK, a)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
	// for _, a := range albumMap {
	// 	if a.ID == id {
	// 		c.IndentedJSON(http.StatusOK, a)
	// 		return
	// 	}
	// }
}
