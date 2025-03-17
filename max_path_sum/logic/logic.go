package logic

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func FindMaxPath(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	} else if len(triangle) == 1 {
		return triangle[0][0]
	}

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += max(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func ReadJsonFile(filePath string) ([][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var triangle [][]int
	err = json.Unmarshal(byteValue, &triangle)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return triangle, nil
}
