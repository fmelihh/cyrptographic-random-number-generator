package main

import (
	"fmt"
	"os"
)

type FileOperator struct {
}

func NewFileOperator() *FileOperator {
	return &FileOperator{}
}

func (f *FileOperator) SaveArrayToTxtFile(strVals []string, fileName string, filePath string) {
	n := len(strVals)
	fullFileName := filePath + "/" + fileName

	file, err := os.Create(fullFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bufferSize := 10_000
	buffer := make([]byte, 0, bufferSize*12)
	for i := 0; i < n; i++ {
		buffer = append(buffer, []byte(strVals[i]+" ")...)
		if len(buffer) >= cap(buffer) {
			_, err := file.Write(buffer)
			if err != nil {
				panic(err)
			}
			buffer = buffer[:0]
		}
	}

	if len(buffer) > 0 {
		_, err := file.Write(buffer)
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("File saved successfully as %s, total: %d", fullFileName, n)
}
