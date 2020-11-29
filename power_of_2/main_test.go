package main

import (
	"interviewbit/helper"
	"testing"
)

func TestPowerOfTwo(t *testing.T) {
	testCases := []struct {
		description    string
		s              string
		expectedResult int
	}{
		{
			description:    "power of two",
			s:              "16",
			expectedResult: 1,
		},
		{
			description:    "not power of two",
			s:              "3",
			expectedResult: 0,
		},
		{
			description:    "big int",
			s:              "147573952589676412928",
			expectedResult: 1,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.description, func(t *testing.T) {
			actualResult := powerOfTwo(testCase.s)

			helper.ShouldBeEqual(t, testCase.expectedResult, actualResult)
		})
	}
}
