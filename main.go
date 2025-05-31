package main

import (
	"fmt"

	"github.com/taylorplewe/md-to-html/md2html"
)

func main() {
	str := md2html.ConvertMarkdownFileToHtml("test.md")
	fmt.Println(str)
}
