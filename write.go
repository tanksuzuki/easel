package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/russross/blackfriday"
)

var cmdWrite = &Command{
	Run:       runWrite,
	UsageLine: "write [OPTIONS] [MARKDOWN]",
	Short:     "Convert the markdown to HTML",
	Long: `
Options:
	-h, --help     Print usage
`,
}

func init() {
}

func runWrite(args []string) int {
	var md []byte

	switch len(args) {
	case 0:
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			md = append(md, []byte(scanner.Text()+"\n")...)
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
	case 1:
		file, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		md = file
	default:
		fmt.Fprintln(os.Stderr, "Too many arguments given.")
		return 1
	}

	m := getCanvasMap(blackfriday.MarkdownCommon(emojify(md)))
	getParsedTemplate().Execute(os.Stdout, m)

	return 0
}
