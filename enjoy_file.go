package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFile() {
	// open file
	inputFile, inputError := os.Open("D:\\pixiv\\test.txt")
	if inputError != nil {
		fmt.Println(inputError)
		os.Exit(1)
	}

	// close file
	defer inputFile.Close()

	// read file
	inputReader := bufio.NewReader(inputFile)
	for {
		text, readError := inputReader.ReadString('\n')
		// text, readError := inputReader.ReadBytes('\n')
		fmt.Println(text)
		if readError == io.EOF {
			return
		}
	}
}

func readCsv() {
	// open file
	inputFile, inputError := os.Open("D:\\pixiv\\test.csv")
	if inputError != nil {
		fmt.Println(inputError)
		os.Exit(1)
	}

	// close file
	defer inputFile.Close()

	var col1, col2 []string
	for {
		var v1, v2 string
		_, err := fmt.Fscanln(inputFile, &v1, &v2)
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
	}
	fmt.Println(col1)
	fmt.Println(col2)
}

func writeFile() {
	buf, readError := ioutil.ReadFile("D:\\pixiv\\test.txt")
	if readError != nil {
		return
	}
	writeError := ioutil.WriteFile("D:\\pixiv\\test1.txt", buf, 0644)
	if writeError != nil {
		return
	}
}

func main() {
	// readFile()
	// readCsv()

	// writeFile()
}
