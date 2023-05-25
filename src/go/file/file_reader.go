package file

import (
    "os"
)

type FileReader struct {}

func NewFileReader() *FileReader {
    return &FileReader{}
}

func (fileReader *FileReader) Read(filename string) (string, error) {
    contents, err := os.ReadFile(filename)
    return string(contents), err
}
