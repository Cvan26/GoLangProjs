package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"ID"`
	Title  string  `json:"Title"`
	Artist string  `json:"Artist"`
	Price  float64 `json:"Price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// write a handler to return all items when client calls api request at /albums, then return all the albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums) // serialize JSON based on the struct album
}

func postAlbum(c *gin.Context) {
	var newalbum album
	if err := c.BindJSON(&newalbum); err != nil {
		return
	}
	albums = append(albums, newalbum)
	c.IndentedJSON(http.StatusCreated, newalbum)
}

func getAlbumById(c *gin.Context) {

	id := c.Param("ID")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
}

func main() {
	router := gin.Default()
	// var c *gin.Context

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.GET("/albums/:ID", getAlbumById)
	router.Run("localhost:8080")
}
