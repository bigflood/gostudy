package store

import (
	"encoding/json"
	"os"
)

func NewInFile(fileName string) *InFile {
	return &InFile{
		fileName: fileName,
	}
}

type InFile struct {
	fileName string
}

func readFromFile(fileName string) (*InMem, error) {
	data := NewInMem()

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return data, nil
	}

	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(f).Decode(data); err != nil {
		return nil, err
	}

	if err := f.Close(); err != nil {
		return nil, err
	}

	return data, nil
}

func writeToFile(fileName string, data *InMem) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	if err := json.NewEncoder(f).Encode(data); err != nil {
		return err
	}

	return f.Close()
}

func (p *InFile) Add(desc string) error {
	data, err := readFromFile(p.fileName)
	if err != nil {
		return err
	}

	if err := data.Add(desc); err != nil {
		return err
	}

	return writeToFile(p.fileName, data)
}

func (p *InFile) List() ([]Task, error) {
	data, err := readFromFile(p.fileName)
	if err != nil {
		return nil, err
	}

	return data.List()
}

func (p *InFile) Done(index int) error {
	data, err := readFromFile(p.fileName)
	if err != nil {
		return err
	}

	if err := data.Done(index); err != nil {
		return err
	}

	return writeToFile(p.fileName, data)
}
