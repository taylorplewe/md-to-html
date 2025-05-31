package md2html

import (
	"strings"
	"unicode"
)

type Converter struct {
	currentLine   string
	start         int
	pos           int
	outBuilder    strings.Builder
	isInCodeBlock bool
}

func (c *Converter) GetHtmlFromMarkdown(input string) string {
	c.outBuilder = strings.Builder{}

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		c.start = 0
		c.pos = 0
		c.currentLine = line
		c.ConvertLine()
	}

	return c.outBuilder.String()
}

func (c *Converter) ConvertLine() {
	for _, r := range c.currentLine {
		if unicode.IsSpace(r) {
			continue
		}

		switch r {
		case '#':
			c.ConvertHeader()
		default:
			c.outBuilder.WriteRune(r)
		}

		c.pos++
	}
}

func (c *Converter) ConvertHeader() {
	for {
		if unicode.IsSpace(s) {

		}
		pos++
	}
}
