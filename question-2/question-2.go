package question_2

import (
	"fmt"
	"math"
)

func isValid(seq []int, encoded string) bool {
	// O(n)
	for i := 0; i < len(encoded); i++ {
		switch encoded[i] {
		case 'L':
			if seq[i] <= seq[i+1] {
				return false
			}
		case 'R':
			if seq[i] >= seq[i+1] {
				return false
			}
		case '=':
			if seq[i] != seq[i+1] {
				return false
			}
		}
	}
	return true
}

func generateSequences(encoded string, seq []int, minSeq *[]int, minSum *int) {
	if len(seq) == len(encoded)+1 {
		if isValid(seq, encoded) {
			sum := 0
			for _, v := range seq {
				sum += v
			}
			if sum < *minSum {
				*minSum = sum
				*minSeq = append([]int(nil), seq...)
			}
		}
		return
	}

	// O(n)
	for i := 0; i <= len(encoded)-1; i++ {
		// O(n)
		generateSequences(encoded, append(seq, i), minSeq, minSum)
	}
}

func AnswerQuestion2(encoded string) {
	minSum := math.MaxInt32
	var minSeq []int

	// O(n^2)
	generateSequences(encoded, []int{}, &minSeq, &minSum)

	output := ""
	for _, num := range minSeq {
		output += fmt.Sprintf("%d", num)
	}

	fmt.Printf("%s: %s\n", encoded, output)
}
