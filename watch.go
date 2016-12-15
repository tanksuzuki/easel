package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jaschaephraim/lrserver"
	"github.com/russross/blackfriday"
	"gopkg.in/fsnotify.v1"
)

var cmdWatch = &Command{
	Run:       runWatch,
	UsageLine: "watch [OPTIONS] MARKDOWN",
	Short:     "Run the preview server",
	Long: `
Options:
	-h, --help     Print usage
	-l, --live     Port for live-reload (Default: 35729)
	-p, --port     Port for preview (Default: 3000)
`,
}

var watchFlags struct {
	port int
	live int
}

func init() {
	cmdWatch.Flag.IntVar(&watchFlags.live, "l", 35729, "")
	cmdWatch.Flag.IntVar(&watchFlags.live, "live", 35729, "")
	cmdWatch.Flag.IntVar(&watchFlags.port, "p", 3000, "")
	cmdWatch.Flag.IntVar(&watchFlags.port, "port", 3000, "")
}

func runWatch(args []string) int {
	switch len(args) {
	case 1:
	case 0:
		fmt.Fprintln(os.Stderr, "\"watch\" requires an argument.")
		return 1
	default:
		fmt.Fprintln(os.Stderr, "Too many arguments given.")
		return 1
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	defer watcher.Close()

	err = watcher.Add(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	lr := lrserver.New(lrserver.DefaultName, uint16(watchFlags.live))
	lr.SetStatusLog(nil)
	lr.SetErrorLog(nil)
	go lr.ListenAndServe()

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				lr.Reload(event.Name)
			case err := <-watcher.Errors:
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}()

	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		md, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		m := getCanvasMap(blackfriday.MarkdownCommon(emojify(md)))
		m["livereload"] = "<script src=\"http://localhost:" + fmt.Sprintf("%d", watchFlags.live) + "/livereload.js\"></script>"
		getParsedTemplate().Execute(rw, m)
	})

	fmt.Printf("listening on localhost:%d\n", watchFlags.port)
	err = http.ListenAndServe(":"+fmt.Sprintf("%d", watchFlags.port), nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	return 0
}
