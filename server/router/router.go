package router

import (
	isprimenot "simple_programs/server/IsPrimeNot"
	"simple_programs/server/palindrome"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.GET("/Palindrome", palindrome.GET_Palindrome)
	router.GET("/IsPrime", isprimenot.Get_IsPrime)
	router.Run("localhost:8081")
}
