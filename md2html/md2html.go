package md2html

import (
	"fmt"
	"io"
	"os"
)

func ConvertMarkdownFileToHtml(filepath string) string {
	inFile, err := os.Open(filepath)
	if err != nil {
		panic(fmt.Errorf("Could not open %s\n", filepath))
	}
	defer inFile.Close()

	inFileBytes, err := io.ReadAll(inFile)
	if err != nil {
		panic(fmt.Errorf("Could not read content of input file\n"))
	}

	inFileString := string(inFileBytes)
	return ConvertMarkdownStringToHtml(inFileString)
}

func ConvertMarkdownStringToHtml(markdownString string) string {
	htmlString := ""

	return htmlString
}
