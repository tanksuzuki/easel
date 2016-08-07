package main

import (
	"fmt"
	"os"
)

const VERSION = "0.2.0"

var cmdVersion = &Command{
	Run:       runVersion,
	UsageLine: "version [OPTIONS]",
	Short:     "Show the easel version information",
	Long: `
Options:
	-h, --help     Print usage
`,
}

func init() {
}

func runVersion(args []string) int {

	if len(args) > 0 {
		fmt.Fprintln(os.Stderr, "Too many arguments given.")
		return 1
	}

	fmt.Printf("easel version %s\n", VERSION)
	return 0
}
