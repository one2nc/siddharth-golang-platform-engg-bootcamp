package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SearchInFile(filePath string, subString string) []string {
	var matchedLines []string
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, subString) {
			matchedLines = append(matchedLines, line)
		}
	}
	return matchedLines
}

func SearchInSTDIN(subString string) []string {
	var inputString []string
	var matchedLines []string
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Err() != nil {
		fmt.Println(scanner.Err().Error())
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "^D" {
			break
		}
		inputString = append(inputString, line)
	}

	for _, input := range inputString {
		if input == subString {
			matchedLines = append(matchedLines, input)
		}
	}

	return matchedLines
}

func WriteOutputToFile(filePath string, output []string) {
	_, error := os.Stat(filePath)

	if os.IsNotExist(error) {
		fo, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer fo.Close()
		for i, _ := range output {
			fo.Write([]byte(output[i]))
		}
	} else {
		fmt.Printf("%v file exist\n", filePath)
		return
	}
}

func main() {
	outputFile := "./output.txt"
	if len(os.Args) != 3 {
		searchString := os.Args[1]
		fmt.Println(SearchInSTDIN(searchString))
		WriteOutputToFile(outputFile, SearchInSTDIN(searchString))
	} else {
		searchString := os.Args[1]
		filePath := os.Args[2]
		fmt.Println(SearchInFile(filePath, searchString))
		WriteOutputToFile(outputFile, SearchInFile(filePath, searchString))
	}
}
