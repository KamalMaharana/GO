package main

// Imports
import (
	"github.com/gin-gonic/gin"
)

// Create as struct of albums
type Albums struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   string `json:"year"`
}

// Create a slice of albums
var albums = []Albums{
	Albums{Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Year: "1973"},
	Albums{Title: "The Wall", Artist: "Pink Floyd", Year: "1979"},
	Albums{Title: "The Final Cut", Artist: "Pink Floyd", Year: "1979"},
	Albums{Title: "Wish You Were Here", Artist: "Pink Floyd", Year: "1975"},
	Albums{Title: "Animals", Artist: "Pink Floyd", Year: "1977"},
	Albums{Title: "The Division Bell", Artist: "Pink Floyd", Year: "1973"},
	Albums{Title: "The Endless River", Artist: "Pink Floyd", Year: "1977"},
	Albums{Title: "The Wall", Artist: "Pink Floyd", Year: "1979"},
	Albums{Title: "The Final Cut", Artist: "Pink Floyd", Year: "1979"},
	Albums{Title: "Wish You Were Here", Artist: "Pink Floyd", Year: "1975"},
	Albums{Title: "Animals", Artist: "Pink Floyd", Year: "1977"},
	Albums{Title: "The Division Bell", Artist: "Pink Floyd", Year: "1973"},
}

// Get all albums
func getAlbums(c *gin.Context) {
	c.JSON(200, albums)
}

// Add a new album
func addAlbum(c *gin.Context) {
	var album Albums
	if err := c.BindJSON(&album); err != nil {
		return
	}
	albums = append(albums, album)
	c.JSON(200, gin.H{"message": "Album added successfully!", "album": album})
}

// Create a new router
func newRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/albums", getAlbums)
	r.POST("/albums", addAlbum)
	return r
}

// Main function
func main() {
	r := newRouter()
	// Listen and Server in
	r.Run(":8080")
}
