package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// ReadLines from file
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// WriteToFile writes content to fileName
func WriteToFile(fileName string, content string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

// ReadFile ...
func ReadFile(file string) ([]byte, error) {
	fileContent, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("util.ReadFile() err  #%v ", err)
	}
	return fileContent, nil
}

func useDefaultEmdOrEmlFileIfInputFileNotSpecified(fileArg string) string {
	inputFile := ""
	if fileExists(fileArg) {
		inputFile = fileArg
	} else {
		if fileExists(defaultEmdFile) {
			inputFile = defaultEmdFile
		} else if fileExists(defaultEmlFile) {
			inputFile = defaultEmlFile
		}
	}
	return inputFile
}

func useDefaultEmdFileIfInputFileNotSpecified(fileArg string) string {
	inputFile := ""
	if fileExists(fileArg) {
		inputFile = fileArg
	} else {
		if fileExists(defaultEmdFile) {
			inputFile = defaultEmdFile
		}
	}
	return inputFile
}

func fileExists(fileName string) bool {
	if _, err := os.Stat(fileName); err == nil {
		return true
	}
	return false
}
