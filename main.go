package main

import (
	"cut/cmd"
)

func main() {
	cli := cmd.NewCmd()
	cli.Execute()
}
