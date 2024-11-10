package cmd

import (
	"cut/parser"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Cmd struct {
	rootCmd *cobra.Command
}

var (
	f int
	d string
)

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
			Args: cobra.ExactArgs(1),
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
	c.rootCmd.PersistentFlags().IntVarP(&f, "field", "f", 0, "The list specifies fields, separated in the input by the field delimiter character (see the -d option). Output fields are separated by a single occurrence of the field delimiter character.")
	c.rootCmd.PersistentFlags().StringVarP(&d, "delimiter", "d", "\t", "Use delim as the field delimiter character instead of the tab character.")

	if err := c.rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
