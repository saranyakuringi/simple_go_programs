package power

import (
	"encoding/json"
	"io"
	"log"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Power struct {
	Num   []int `json:"num"`
	Power int   `json:"power"`
}

func Power_Of_the_number(num []int, power int) (results []float64) {
	//var results []float64
	for i := 0; i < len(num); i++ {
		result := math.Pow(float64(num[i]), float64(power))
		results = append(results, result)
	}
	return results
}

func GET_Power_Of_the_number(c *gin.Context) {
	request, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error in request body", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error in request body"})
		return
	}
	var data Power
	err = json.Unmarshal(request, &data)
	if err != nil {
		log.Println("Error in unmarshal", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error in unmarshal"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"Power of the number is:": Power_Of_the_number(data.Num, data.Power)})
}
