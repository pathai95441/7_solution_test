package main

import (
	"fmt"
	"log"

	"max_path_sum/logic"
	"max_path_sum/validate"
)

func main() {
	triangle, err := logic.ReadJsonFile("./test_case/case_1.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	if !validate.IsValidTriangle(triangle) {
		log.Fatal("Invalid Triangle")
		return
	}

	maxSum := logic.FindMaxPath(triangle)
	fmt.Println("Maximum Path Sum:", maxSum)
}
