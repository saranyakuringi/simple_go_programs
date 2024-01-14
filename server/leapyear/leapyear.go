package leapyear

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Year struct {
	Year []int `json:"year"`
}

func leapyear(year []int) (leapyear []int, Notleapyear []int) {

	for _, value := range year {
		if value%4 == 0 {
			leapyear = append(leapyear, value)

		} else {
			Notleapyear = append(Notleapyear, value)
		}
	}
	return leapyear, Notleapyear
}

func GET_leapyear(c *gin.Context) {
	request, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error in request body", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error in request body"})
		return
	}
	var year Year
	err = json.Unmarshal(request, &year)
	if err != nil {
		log.Println("Error in unmarshal", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error in unmarshal"})
		return
	}
	leapyears, Notleapyears := leapyear(year.Year)
	c.IndentedJSON(http.StatusOK, gin.H{"Leap Year": leapyears, "Not a Leap Year": Notleapyears})
}
