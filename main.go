package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type note struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Content  string `json:"content"`
}

var notes = []note{
	{ID: "1", Title: "Kill The President", Subtitle: "Worse than JFK", Content: "lirem epsum cywbcybrbvhbev chebcwebcw cebcwycbe"},
	{ID: "2", Title: "Kill The President", Subtitle: "Worse than JFK", Content: "lirem epsum cywbcybrbvhbev chebcwebcw cebcwycbe"},
	{ID: "3", Title: "Kill The President", Subtitle: "Worse than JFK", Content: "lirem epsum cywbcybrbvhbev chebcwebcw cebcwycbe"},
}

func main() {
	router := gin.Default()
	router.GET("/notes", getNotes)
	router.POST("/notes", postNote)
	router.GET("/notes/:id", getNoteByID)

	router.Run("localhost:8080")
}

func getNotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, notes)
}

func postNote(c *gin.Context) {
	var newNote note

	if err := c.BindJSON(&newNote); err != nil {
		return
	}

	notes = append(notes, newNote)
	c.IndentedJSON(http.StatusCreated, newNote)
}

func getNoteByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range notes {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "note not found"})
}
