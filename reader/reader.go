package reader

import (
	"io"
	"log"
	"os"
	"errors"
)

func ReadFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	failed := "Not possible read file"

	return fileBytes, errors.New(failed)
}
