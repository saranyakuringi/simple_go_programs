package stringreversal

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//define type

type Str_reversal struct {
	Word string `json:"word"`
}

//define string reversal function

func String_Reversal(s string) string {
	//rns := []rune(s)
	words := strings.Fields(s)

	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}

	return strings.Join(words, " ")
}

// define function to read data from postman and display the output
func GET_String_Reversal(c *gin.Context) {
	request, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error in request body", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error in request body"})
		return
	}
	var result Str_reversal
	err = json.Unmarshal(request, &result)
	if err != nil {
		log.Println("Error in unmarshal", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error in unmarshal"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"Original String": result, "The Reversed string": String_Reversal(result.Word)})
}
