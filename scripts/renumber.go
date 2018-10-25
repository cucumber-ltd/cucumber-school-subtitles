package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	var files []string

	if len(os.Args) > 1 && os.Args[1] != "" {
		absFile, err := filepath.Abs(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		files = []string{absFile}
	} else {
		fmt.Println("no file provided, processing directory")
		files = directoryVTTFiles()
	}

	err := processFiles(files)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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
		fileData, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("unable to open file %s error: %v", file, err)
		}
		defer fileData.Close()

		// renumber the file
		renumberedFile, err := renumberFile(fileData)
		if err != nil {
			return err
		}

		// write the file
		fmt.Println("Writing to file " + filepath.Base(file))
		if err := ioutil.WriteFile(file, renumberedFile, 0777); err != nil {
			return fmt.Errorf("unable to write to file %v", err)
		}

		return nil
	}

	return nil
}

func renumberFile(r io.Reader) ([]byte, error) {
	scanner := bufio.NewScanner(r)
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

func toByteSlice(in interface{}) []byte {
	return []byte(fmt.Sprintf("%v\n", in))
}
