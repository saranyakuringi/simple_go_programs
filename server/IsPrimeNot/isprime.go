package isprimenot

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Prime_type struct {
	Num []int `json:"num"`
}

// function to check if the numbers in the slice are prime or not
func IsPrime(n []int) (primes []int, notprime []int) {
	for _, num := range n {
		if num <= 1 {
			notprime = append(notprime, num)
		} else {
			prime := true
			for j := 2; j < num; j++ {
				if num%j == 0 {
					prime = false
					//notprime = append(notprime, num)
					break
				}
			}
			if prime {
				primes = append(primes, num)
			} else {
				notprime = append(notprime, num)
			}
		}
	}
	return primes, notprime
}

// function to read data from postman and display the output
func Get_IsPrime(c *gin.Context) {
	request, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error in request body", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error in request body"})
		return
	}
	var Num Prime_type
	err = json.Unmarshal(request, &Num)
	if err != nil {
		log.Println("Error in unmarshal", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error in unmarshal"})
		return
	}
	prime, NotPrime := IsPrime(Num.Num)
	c.IndentedJSON(http.StatusOK, gin.H{"Prime": prime, "Not Prime": NotPrime})
}
