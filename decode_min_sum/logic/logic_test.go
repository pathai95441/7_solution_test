package logic

import (
	"testing"
)

type TestArgs struct {
	input    string
	expected string
}

func TestDecode(t *testing.T) {
	test_case := []TestArgs{
		{input: "LLRR=", expected: "210122"},
		{input: "==RLL", expected: "000210"},
		{input: "=LLRR", expected: "221012"},
		{input: "RRL=R", expected: "012001"},
	}

	for _, tc := range test_case {
		result := DecodeMinSum(tc.input)

		if result != tc.expected {
			t.Errorf("For input %s, expected %s, but got %s", tc.input, tc.expected, result)
		}
	}
}
