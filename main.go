package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"max/backend/albums/models"
)

func getAlbums(c *gin.Context) {
	data := models.Get_albums_from_db()
	c.IndentedJSON(http.StatusOK, data)
}

func getAlbumWidthId(c *gin.Context) {
	idP := c.Param("id")
	id,_ := strconv.Atoi(idP)
	data := models.Get_albums_from_db_width_id(id)
	c.IndentedJSON(http.StatusOK, data)
}

// postAlbum...
func postAlbums(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	models.Insert_to_db(&newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	models.Db_init()
	defer models.Db.Close()
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumWidthId)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
