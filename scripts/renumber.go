package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	files := filesToProcess()
	err := processFiles(files)
	if err != nil {
		panic(err)
	}
}

func filesToProcess() []string {
	if len(os.Args) > 1 && os.Args[1] != "" {
		return []string{os.Args[1]}
	}

	return directoryVTTFiles()
}

func directoryVTTFiles() []string {
	searchDir, _ := os.Getwd()

	var fileList []string
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if strings.Contains(path, ".vtt") {
			fileList = append(fileList, path)
		}
		return nil
	})
	if err != nil {
		fmt.Print("err:", err)
	}

	return fileList
}

func processFiles(files []string) error {
	for _, file := range files {
		// read the file
		fileData, err := read(file)
		if err != nil {
			return err
		}
		defer fileData.Close()

		// renumber the file
		renumberedFile, err := renumberFile(fileData)
		if err != nil {
			return err
		}

		// write the file
		fmt.Println("Writing to:", file)
		write(file, renumberedFile)
	}

	return nil
}

func renumberFile(fileData *os.File) ([]byte, error) {
	scanner := bufio.NewScanner(fileData)
	scanner.Split(bufio.ScanLines)

	frameTracker := 1
	fileOutData := []byte{}
	for scanner.Scan() {
		line := scanner.Text()

		// if the line asserts to an int, we know it's a frame number,
		// otherwise just add it to the byte slice
		_, err := strconv.Atoi(line)
		if err == nil {
			fileOutData = append(fileOutData, toByteSlice(frameTracker)...)
			frameTracker++
		} else {
			fileOutData = append(fileOutData, toByteSlice(line)...)
		}
	}

	return fileOutData, nil
}

func read(fileLocation string) (*os.File, error) {
	fileData, err := os.Open(fileLocation)
	if err != nil {
		return nil, fmt.Errorf("unable to open file %s error: %v", fileLocation, err)
	}

	return fileData, nil
}

func write(fileLocation string, data []byte) error {
	if err := ioutil.WriteFile(fileLocation, data, 0777); err != nil {
		return fmt.Errorf("unable to write to file %v", err)
	}
	return nil
}

func toByteSlice(in interface{}) []byte {
	return []byte(fmt.Sprintf("%v\n", in))
}
