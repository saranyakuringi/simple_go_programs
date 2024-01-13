package router

import (
	isprimenot "simple_programs/server/IsPrimeNot"
	"simple_programs/server/palindrome"
	"simple_programs/server/stringreversal"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.GET("/Palindrome", palindrome.GET_Palindrome)
	router.GET("/IsPrime", isprimenot.Get_IsPrime)
	router.GET("/StringReversal", stringreversal.GET_String_Reversal)
	router.Run("localhost:8081")
}
