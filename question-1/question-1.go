package question_1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func maxPathSum(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += findMax(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func AnswerQuestion1() {
	triangle := [][]int{{59}, {73, 41}, {52, 40, 53}, {26, 53, 6, 34}}
	fmt.Println("Test Case {{59}, {73, 41}, {52, 40, 53}, {26, 53, 6, 34}}: ", maxPathSum(triangle))

	file, err := ioutil.ReadFile("files/hard.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	var jsonData [][]int
	err = json.Unmarshal(file, &jsonData)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		os.Exit(1)
	}

	fmt.Println("JSON Input:", maxPathSum(jsonData))
}
