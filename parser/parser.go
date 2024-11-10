package parser

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type Parser struct {
	path       string
	fieldCount int
	delimeter  string
}

func NewParser(path string, f int) *Parser {
	return &Parser{
		path:       path,
		fieldCount: f,
		delimeter:  "\t",
	}
}

func (p *Parser) Parse() (string, error) {
	file, err := os.Open(p.path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return p.FindFields(file)
}

func (p *Parser) FindFields(reader io.Reader) (string, error) {
	scanner := bufio.NewScanner(reader)
	var parts []string

	// Process the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue // Skip empty lines
		}

		part := p.FindNthField(line)
		parts = append(parts, part)
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strings.Join(parts, "\n"), nil
}

func (p *Parser) FindNthField(line string) string {
	lineParts := strings.Split(line, p.delimeter)

	// Handle index out of range
	if p.fieldCount > len(lineParts) {
		return ""
	}

	return lineParts[p.fieldCount-1]
}
