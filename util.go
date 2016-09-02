package main

import (
	"regexp"
	"strings"
	"text/template"
)

func getTitle(html []byte) string {
	r := regexp.MustCompile("<h1>(.*)</h1>")
	return strings.Replace(r.FindStringSubmatch(string(html))[1], " ", "", -1)
}

func getContent(html []byte) string {
	r := regexp.MustCompile("<h1>(.*)</h1>")
	return strings.TrimSpace(r.ReplaceAllString(string(html), ""))
}

func getCanvasMap(html []byte) map[string]string {
	r := regexp.MustCompile("<h1>(\\w|\\s)*</h1>")
	index := r.FindAllIndex(html, -1)

	m := map[string]string{}
	for i := 0; i < len(index); i++ {
		if i == (len(index) - 1) {
			m[getTitle(html[index[i][0]:])] = getContent(html[index[i][0]:])
		} else {
			m[getTitle(html[index[i][0]:index[i+1][0]])] = getContent(html[index[i][0]:index[i+1][0]])
		}
	}

	return m
}

func getParsedTemplate() *template.Template {
	tpl := template.New("")
	tpl = template.Must(tpl.Parse(TEMPLATE_CANVAS))
	tpl = template.Must(tpl.Parse(TEMPLATE_SANITIZE))
	tpl = template.Must(tpl.Parse(TEMPLATE_STYLE_CANVAS))
	tpl = template.Must(tpl.Parse(TEMPLATE_STYLE_MARKDOWN))
	return tpl
}
