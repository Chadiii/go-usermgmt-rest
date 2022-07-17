package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []user{
	{Name: "Alice", Age: 43, Id: 81}, {Name: "Bob", Age: 30, Id: 887}, {Name: "Alice", Age: 43, Id: 81}, {Name: "Bob", Age: 30, Id: 887},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func postAlbums(c *gin.Context) {
	var newUser user

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	// Add the new album to the slice.
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
