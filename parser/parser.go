package parser

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Parser struct {
	path       string
	fieldCount []int
	delimeter  string
}

func NewParser(path string, f []int, d string) *Parser {
	return &Parser{
		path:       path,
		fieldCount: f,
		delimeter:  d,
	}
}

func (p *Parser) Parse() error {
	reader, err := p.GetReader()
	if err != nil {
		return err
	}
	defer reader.Close()

	return p.FindFields(reader, os.Stdout)
}

func (p *Parser) GetReader() (*os.File, error) {
	if p.path == "" {
		return os.Stdin, nil
	}

	file, err := os.Open(p.path)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (p *Parser) FindFields(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)

	// Process the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue // Skip empty lines
		}

		part := p.FindNthField(line)
		if part != "" {
			// Write each line immediately
			if _, err := fmt.Fprintln(writer, part); err != nil {
				return err
			}
		}
	}

	return scanner.Err()
}

func (p *Parser) FindNthField(line string) string {
	lineParts := strings.Split(line, p.delimeter)

	str := ""

	// Handle index out of range
	for i := range p.fieldCount {
		if p.fieldCount[i] > len(lineParts) {
			return ""
		}

		str = str + "," + lineParts[p.fieldCount[i]-1]
	}

	return strings.TrimLeft(str, ",")
}
