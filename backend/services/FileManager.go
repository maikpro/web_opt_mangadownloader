package services

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func SaveFile(path string, filename string, file []byte) (string, error) {
	fullpath := fmt.Sprintf("%s/%s", path, filename)

	// Create the directory if it doesn't exist
	err := os.MkdirAll(filepath.Dir(fullpath), 0o777)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// Write the data to a file
	err = os.WriteFile(fullpath, file, 0o777)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return filepath.Abs(path)
}

func CreateZip(filesDirectory string) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	dir := strings.ReplaceAll(filesDirectory, "\\", "/")

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		// Skip directories
		if info.IsDir() {
			log.Printf("skipping a dir without errors: %+v \n", info.Name())
			return nil
		}

		// Open the image file
		file, err := os.Open(path)
		if err != nil {
			log.Println("failed to open file")
			return err
		}
		defer file.Close()

		// Create a new zip file entry
		zipFile, err := zipWriter.Create(info.Name())
		if err != nil {
			log.Println("failed to create file in zip")
			return err
		}

		// Copy the file contents into the zip file entry
		_, err = io.Copy(zipFile, file)
		if err != nil {
			log.Println("failed to copy file contents to zip")
			return err
		}

		return nil
	})

	// Close the zip writer to finalize the zip file
	zipWriter.Close()
	return buf, err
}
