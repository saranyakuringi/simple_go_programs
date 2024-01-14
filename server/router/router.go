package router

import (
	isprimenot "simple_programs/server/IsPrimeNot"
	"simple_programs/server/leapyear"
	"simple_programs/server/palindrome"
	"simple_programs/server/power"
	"simple_programs/server/stringreversal"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.GET("/Palindrome", palindrome.GET_Palindrome)
	router.GET("/IsPrime", isprimenot.Get_IsPrime)
	router.GET("/StringReversal", stringreversal.GET_String_Reversal)
	router.GET("/Isleap", leapyear.GET_leapyear)
	router.GET("/power", power.GET_Power_Of_the_number)
	router.Run("localhost:8081")
}
