package palindrome

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// define structure
type String_type struct {
	Word string `json:"word"`
}

// Function to check palindrome
func Palindrome(name string) bool {
	//name = "kayak"
	var data bool
	fmt.Println(len(name))
	for i := 0; i < len(name)/2; i++ {
		if name[i] == name[len(name)-i-1] {
			data = true
			continue
		} else {
			data = false
			break
		}
	}
	//fmt.Println("This is a palindrome")
	return data
}

// function to read data from postman and display the output
func GET_Palindrome(c *gin.Context) {
	request, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error in request body", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error in request body"})
		return
	}

	var word String_type
	err = json.Unmarshal(request, &word)
	if err != nil {
		log.Println("Error in unmarshal", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error in unmarshal"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"Is is Palindrome: ": Palindrome(word.Word)})
}
