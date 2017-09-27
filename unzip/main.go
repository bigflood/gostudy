package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func main() {

	if len(os.Args) != 3 {
		panic("usage unzip input_zip_file dest_dir")
	}

	inputFile := os.Args[1]
	outputDir := os.Args[2]

	err := unzipFile(inputFile, outputDir)
	if err != nil {
		panic(err)
	}
}

func unzipFile(inputFile, outputDir string) error {

	const numRt = 2
	fmt.Println("numRt:", numRt)

	startTime := time.Now()

	r, err := zip.OpenReader(inputFile)
	if err != nil {
		return err
	}

	defer r.Close()

	resultChan := make(chan int64)
	limit := make(chan int, numRt)
	errChan := make(chan error)

	fileCount := 0

	for _, f := range r.File {
		if f.UncompressedSize64 == 0 {
			continue
		}

		fileCount++

		go func(f *zip.File) {
			limit <- 0
			defer func() { <-limit }()

			rc, err := f.Open()
			if err != nil {
				errChan <- err
				return
			}

			w, err := decompressFile(rc, filepath.Join(outputDir, f.Name))
			if err != nil {
				errChan <- err
				return
			}

			rc.Close()

			resultChan <- w
		}(f)
	}

	processCount := 0
	var numBytes int64

MainLoop:
	for {
		select {
		case w := <-resultChan:
			processCount++
			numBytes += w
			fmt.Print(processCount, "/", fileCount, " ", numBytes, " bytes\r")
			if processCount == fileCount {
				break MainLoop
			}
		case err := <-errChan:
			panic(err)
		}
	}

	fmt.Println()
	fmt.Println(time.Since(startTime))

	return nil
}

func decompressFile(input io.Reader, outputFileName string) (int64, error) {
	err := os.MkdirAll(filepath.Dir(outputFileName), os.ModeDir)
	if err != nil {
		return 0, err
	}

	file, err := os.Create(outputFileName)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	bytesWritten, err := io.Copy(file, input)
	return bytesWritten, err
}
