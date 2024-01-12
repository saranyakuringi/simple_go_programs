package router

import (
	"simple_programs/server/palindrome"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.GET("/Palindrome", palindrome.GET_Palindrome)
	router.Run("localhost:8081")
}
