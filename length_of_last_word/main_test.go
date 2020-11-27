package main

import (
	"testing"

	"interviewbit/helper"
)

func TestLengthOfLastWord(t *testing.T) {
	testCases := []struct {
		description    string
		s              string
		expectedLength int
	}{
		{
			description:    "return last length - single word",
			s:              "Hello",
			expectedLength: 5,
		},
		{
			description:    "return last length - multiple word",
			s:              "It's a great day",
			expectedLength: 3,
		},
		{
			description:    "empty string",
			s:              "",
			expectedLength: 0,
		},
		{
			description:    "space character",
			s:              " ",
			expectedLength: 0,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.description, func(t *testing.T) {
			actualLength := LengthOfLastWord(testCase.s)

			helper.ShouldBeEqual(t, testCase.expectedLength, actualLength)
		})
	}
}
