package main

import (
	"fmt"
	"simple_programs/server/router"
)

func main() {
	fmt.Println("This is an API for simple programs")
	router.Router()

}

// func Power_Of_the_number() {
// 	num := []int{1, 2, 3}
// 	power := 3
// 	var results []float64
// 	for i := 0; i < len(num); i++ {
// 		result := math.Pow(float64(num[i]), float64(power))
// 		results = append(results, result)
// 	}
// 	fmt.Println("Results", results)
// }
