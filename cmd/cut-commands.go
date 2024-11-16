package cmd

import (
	"cut/parser"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Cmd struct {
	rootCmd *cobra.Command
}

var (
	f    []int
	d    string
	fStr string
)

// parseFields converts a string of numbers (comma or space separated) into a slice of integers
func parseFields(s string) ([]int, error) {
	// Handle empty input
	if len(strings.TrimSpace(s)) == 0 {
		return nil, nil
	}

	// First split by comma
	parts := strings.Split(s, ",")

	var result []int
	for _, part := range parts {
		// Then split by whitespace
		numbers := strings.Fields(strings.TrimSpace(part))
		for _, numStr := range numbers {
			num, err := strconv.Atoi(strings.TrimSpace(numStr))
			if err != nil {
				return nil, fmt.Errorf("invalid field number: %s", numStr)
			}
			if num < 1 {
				return nil, fmt.Errorf("field numbers must be positive: %d", num)
			}
			result = append(result, num)
		}
	}

	return result, nil
}

func NewCmd() *Cmd {
	return &Cmd{
		rootCmd: &cobra.Command{
			Use:   "cut",
			Short: "cut out selected portions of each line of a file",
			Long: `The cut utility cuts out selected portions of each line (as specified by list) from each file and writes them to
				     the standard output.  If no file arguments are specified, or a file argument is a single dash (‘-’), cut reads
				     from the standard input.  The items specified by list can be in terms of column position or in terms of fields
				     delimited by a special character.  Column and field numbering start from 1.

				     The list option argument is a comma or whitespace separated set of numbers and/or number ranges.  Number ranges
				     consist of a number, a dash (‘-’), and a second number and select the columns or fields from the first number to
				     the second, inclusive.  Numbers or number ranges may be preceded by a dash, which selects all columns or fields
				     from 1 to the last number.  Numbers or number ranges may be followed by a dash, which selects all columns or
				     fields from the last number to the end of the line.  Numbers and number ranges may be repeated, overlapping, and
				     in any order.  If a field or column is specified multiple times, it will appear only once in the output.  It is
				     not an error to select columns or fields not present in the input line.
				     `,
			PreRunE: func(cmd *cobra.Command, args []string) error {
				var err error
				f, err = parseFields(fStr)
				if err != nil {
					return err
				}
				return nil
			},
			Run: func(cmd *cobra.Command, args []string) {
				p := parser.NewParser(args[0], f, d)
				err := p.Parse()
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}
				os.Exit(0)
			},
		},
	}
}

func (c *Cmd) Execute() {
	c.rootCmd.PersistentFlags().StringVarP(&fStr, "field", "f", "",
		"The list specifies fields, separated by commas or spaces (e.g., '1,2,3' or '1 2 3')")
	c.rootCmd.PersistentFlags().StringVarP(&d, "delimiter", "d", "\t", "Use delim as the field delimiter character instead of the tab character.")

	if err := c.rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
