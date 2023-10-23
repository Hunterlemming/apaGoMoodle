package docx

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func Unzip(docName string) (string, error) {
	reader, err := zip.OpenReader(docName)
	if err != nil {
		fmt.Println("Error opening .docx file:", err)
		return "", err
	}
	defer reader.Close()

	os.Mkdir(DocExtractDir, os.ModeDir)
	resultName := ""

	for _, file := range reader.File {
		if file.Name != "word/document.xml" {
			continue
		}

		fileViewer, err := file.Open()
		if err != nil {
			fmt.Println("Error opening file in .docx:", err)
			return "", err
		}
		defer fileViewer.Close()

		resultName = fmt.Sprintf("%s/input.xnl", DocExtractDir)
		extractFile, err := os.Create(resultName)
		if err != nil {
			fmt.Println("Error creating extracted file:", err)
			return "", err
		}
		defer extractFile.Close()

		_, err = io.Copy(extractFile, fileViewer)
		if err != nil {
			fmt.Println("Error extracting file:", err)
			return "", err
		}
	}

	return resultName, nil
}

func Cleanup() {
	err := os.RemoveAll(DocExtractDir)
	if err != nil {
		fmt.Println("Error deleting directory:", err)
		return
	}
}
