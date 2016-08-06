package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/russross/blackfriday"
)

var cmdWrite = &Command{
	Run:       runWrite,
	UsageLine: "write [OPTIONS] MARKDOWN",
	Short:     "Convert the markdown to HTML",
	Long: `
Options:
	-h, --help     Print usage
`,
}

func init() {
}

func runWrite(args []string) int {
	switch {
	case len(args) == 1:
	case len(args) == 0:
		fmt.Fprintln(os.Stderr, "\"write\" requires an argument.")
		return 1
	default:
		fmt.Fprintln(os.Stderr, "Too many arguments given.")
		return 1
	}

	md, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	m := getCanvasMap(blackfriday.MarkdownCommon(md))

	tpl := template.New("")
	template.Must(tpl.Parse(HTML_TEMPLATE))
	tpl.Execute(os.Stdout, m)

	return 0
}
