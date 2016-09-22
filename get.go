package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/russross/blackfriday"
)

var cmdGet = &Command{
	Run:       runGet,
	UsageLine: "get [OPTIONS] URL",
	Short:     "Preview the markdown on the web",
	Long: `
Options:
	-h, --help     Print usage
	-p, --port     Port for preview (Default: 3000)
`,
}

var getFlags struct {
	port int
}

func init() {
	cmdGet.Flag.IntVar(&getFlags.port, "p", 3000, "")
	cmdGet.Flag.IntVar(&getFlags.port, "port", 3000, "")
}

func runGet(args []string) int {
	switch len(args) {
	case 1:
	case 0:
		fmt.Fprintln(os.Stderr, "\"get\" requires an argument.")
		return 1
	default:
		fmt.Fprintln(os.Stderr, "Too many arguments given.")
		return 1
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		resp, err := http.Get(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		defer resp.Body.Close()

		md, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		m := getCanvasMap(blackfriday.MarkdownCommon(emojify(md)))
		getParsedTemplate().Execute(rw, m)
	})

	fmt.Printf("listening on localhost:%d\n", getFlags.port)
	err := http.ListenAndServe(":"+fmt.Sprintf("%d", getFlags.port), nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	return 0
}
