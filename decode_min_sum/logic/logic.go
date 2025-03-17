package logic

import "fmt"

func DecodeMinSum(encoded string) string {
	n := len(encoded) + 1
	seq := make([]int, n)
	seq[0] = 0
	countR := 1

	for i, ch := range encoded {
		switch ch {
		case 'L':
			if i-1 >= 0 && encoded[i-1] == 'L' {
				seq = sumDecodedLogicL(i, seq, encoded)
			} else {
				seq[i+1] = 0
			}
		case 'R':
			seq[i+1] = countR
			if i+1 < n-1 && encoded[i+1] == 'R' {
				countR++
			} else {
				countR = 1
			}

		case '=':
			seq[i+1] = seq[i]
		}
	}

	result := ""
	for _, num := range seq {
		result += fmt.Sprintf("%d", num)
	}
	return result
}

func sumDecodedLogicL(c int, seq []int, encoded string) []int {
	countL := 1
	seq[c+1] = 0
	for i := c; i >= 0; i-- {
		current := encoded[i]
		if current == '=' {
			return replaceEqual(i, seq, encoded, seq[i+1])
		} else if current == 'R' {
			return seq
		}
		seq[i] = countL
		countL++
	}
	return seq
}

func replaceEqual(c int, seq []int, encoded string, equal int) []int {
	for i := c; i >= 0; i-- {
		current := encoded[i]
		if current != '=' {
			return seq
		}
		seq[i] = equal
	}
	return seq
}
