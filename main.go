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
	router.GET("/albums", getAlbums)
	router.GET("/albums:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}

// postAlbums adds an album from JSON received in the request body
func postAlbums(c *gin.Context){
	var newAlbum album

// call BindJSON to bind the received Json to newAlbum
if err := c.BindJSON(&newAlbum); err != nil {
	return
}

// add the new albulm to the slice 
albums = append(albums ,newAlbum)
c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context){
	id := c.Param("id")

	// loop over the list of albums looking for an album  whose id value matches the parameter
	for _, a := range albums{
		if a.ID == id{
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"albulm not found!!"})
}

