package store

import (
	"encoding/json"
)


func NewInFile(fileName string) *InFile {
	return &InFile{
		file: NewFile(fileName),
	}
}

func NewFromDataSource(file DataSource) *InFile {
	return &InFile{
		file: file,
	}
}

type InFile struct {
	file DataSource
}

func readFromFile(file DataSource) (*InMem, error) {
	data := NewInMem()

	bytes, err := file.ReadAll()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, data); err != nil {
		return nil, err
	}

	return data, nil
}

func writeToFile(file DataSource, data *InMem) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return file.WriteAll(b)
}

func (p *InFile) Add(desc string) error {
	data, err := readFromFile(p.file)
	if err != nil {
		return err
	}

	if err := data.Add(desc); err != nil {
		return err
	}

	return writeToFile(p.file, data)
}

func (p *InFile) List(filter Filter) ([]Task, error) {
	data, err := readFromFile(p.file)
	if err != nil {
		return nil, err
	}

	return data.List(filter)
}

func (p *InFile) Done(index int) error {
	data, err := readFromFile(p.file)
	if err != nil {
		return err
	}

	if err := data.Done(index); err != nil {
		return err
	}

	return writeToFile(p.file, data)
}
