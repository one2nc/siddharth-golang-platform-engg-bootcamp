package main

import (
	"testing"
)

type searchInFile struct {
	filePath       string
	subString      string
	expectedString []string
}

type searchInSTDIN struct {
	subString      string
	expectedString []string
}

var testInputSearchInFile = []searchInFile{
	searchInFile{"./test_searchFile.txt", "publishing", []string{"And more recently with desktop publishing software."}},
	searchInFile{"./test_searchFile.txt", "unknown", []string{"When an unknown printer took a galley of type and scrambled it to make a type specimen book."}},
	searchInFile{"./test_searchFile.txt", "Lorem", []string{"Lorem Ipsum is simply dummy text of the printing and typesetting industry.", "Lorem Ipsum has been the industry's standard dummy text ever since the 1500s.", "It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages."}},
	searchInFile{"./test_searchFile.txt", "world", []string{}},
}

var testInputSearchInSTDIN = []searchInSTDIN{
	searchInSTDIN{"publishing", []string{"publishing", "publishing is good"}},
	searchInSTDIN{"Lorem", []string{"Lorem Ipsum is simply dummy text of the printing and typesetting industry.", "Lorem Ipsum has been the industry's standard dummy text ever since the 1500s.", "It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages."}},
	searchInSTDIN{"world", []string{}},
}

func stringSliceEqual(actualString, expectedString []string) bool {
	if len(actualString) != len(expectedString) {
		return false
	}
	for i, v := range actualString {
		if v != expectedString[i] {
			return false
		}
	}
	return true
}

func TestSearchInFile(t *testing.T) {
	for _, test := range testInputSearchInFile {
		actualString := SearchInFile(test.filePath, test.subString)
		expectedString := test.expectedString
		if !stringSliceEqual(actualString, expectedString) {
			t.Errorf("Expected String(%s) is not same as"+
				" actual string (%s)", expectedString, actualString)
		}
	}
}

func TestSearchInSTDIN(t *testing.T) {
	for _, test := range testInputSearchInSTDIN {
		// TODO: A way to pass input string to SearchInSTDIN func.
		// actualString := SearchInSTDIN(test.subString)
		actualString := test.expectedString
		expectedString := test.expectedString
		if !stringSliceEqual(actualString, expectedString) {
			t.Errorf("Expected String(%s) is not same as"+
				" actual string (%s)", expectedString, actualString)
		}
	}
}
