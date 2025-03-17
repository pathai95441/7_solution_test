package logic

import (
	"testing"
)

type TestArgs struct {
	input    string
	expected int
}

func TestFindMaxPath(t *testing.T) {
	test_case := []TestArgs{
		{input: "./test_case/case_1.json", expected: 237},
		{input: "./test_case/case_2.json", expected: 7273},
	}

	for _, tc := range test_case {
		triangle, err := ReadJsonFile(tc.input)
		if err != nil {
			t.Error(err)
		}

		result := FindMaxPath(triangle)

		if result != tc.expected {
			t.Errorf("Expected %d, but got %d", tc.expected, result)
		}

	}
}
