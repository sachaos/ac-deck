package tester

import (
	"fmt"
	"testing"
)

func Test_judgeEquality(t *testing.T) {
	testCases := []struct{
		example string
		actual string
		expected bool
	}{
		{
			example:  "0",
			actual:   "0",
			expected: true,
		},
		{
			example:  "hoge",
			actual:   "hogehoge",
			expected: false,
		},
		{
			example: "3000293",
			actual: "3000292",
			expected: false,
		},
		{
			example:  "100",
			actual:   "100",
			expected: true,
		},
		{
			example: "2.6666666667",
			actual:  "2.6666666666666665",
			expected: true,
		},
		{
			example: "2.6666666667",
			actual:  "2.6667",
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%s_%s", testCase.example, testCase.actual), func(t *testing.T) {
			if result := judgeEquality(testCase.example, testCase.actual); result != testCase.expected {
				t.Errorf("expected %v, but actually %v", testCase.expected, result)
			}
		})
	}
}
