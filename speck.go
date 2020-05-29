package main

import (
	"bytes"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	TagName = "speck"
	AttrTab = "tab"
)

func main() {
	err := PrintFileSections(os.Args[1:], os.Stdout)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// PrintFileSections iterates over a list of files,
// finds Speck-specific tag sections, and prints them to writer
func PrintFileSections(files []string, out io.Writer) error {
	for _, filename := range files {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
		if err != nil {
			return err
		}
		var innerErr error
		doc.Find(TagName).Each(func(i int, el *goquery.Selection) {
			if innerErr != nil {
				return
			}
			contents, err := getElementContents(el)
			if err != nil {
				innerErr = err
				return
			}
			fmt.Fprintln(out, contents)
		})
		if innerErr != nil {
			return innerErr
		}
	}
	return nil
}

func getElementContents(el *goquery.Selection) (string, error) {
	contents, err := el.Html()
	if err != nil {
		return "", err
	}
	contents = html.UnescapeString(contents)
	contents = strings.TrimPrefix(contents, "\n")
	tabsStr := el.AttrOr(AttrTab, "0")
	tabs, err := strconv.Atoi(tabsStr)
	if err != nil {
		return "", err
	}
	if tabs > 0 {
		contents = deTabulate(contents, tabs)
	}
	return contents, nil
}

func deTabulate(contents string, tabs int) string {
	prefix := strings.Repeat("\t", tabs)
	var lines []string
	for _, line := range strings.Split(contents, "\n") {
		lines = append(lines, strings.TrimPrefix(line, prefix))
	}
	contents = strings.Join(lines, "\n")
	return contents
}
