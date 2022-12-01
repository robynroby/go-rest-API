package main


import(
	"net/http"
	"github.com/gin-gonic/gin"
)

// albulm represents data about a record albulm
type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

// albulm slice to seed record albulm data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Lengend", Price: 54.89},
	{ID: "2", Title: "Sarah", Artist: "Jonnas trnet", Price: 34.89},
	{ID: "3", Title: "Peru", Artist: "Rema", Price: 64.09},
}

// getAlbums responds with the list of all albums as jSON
func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums" , getAlbums)
	router.Run("localhost:8080")
}