package store

import "io/ioutil"

type DataSource interface {
	ReadAll() ([]byte, error)
	WriteAll([]byte) error
}

//go:generate counterfeiter -o mocks/file.go --fake-name File . DataSource

func NewFile(fileName string) *File {

	return &File {
		fileName:fileName,
	}
}

type File struct {
	fileName string
}

func (file *File) ReadAll() ([]byte, error) {
	return ioutil.ReadFile(file.fileName)
}

func (file *File) WriteAll(data []byte) error {
	return ioutil.WriteFile(file.fileName, data, 0666)
}
