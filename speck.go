package main

import (
	"bytes"
	"fmt"
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
		doc.Find(TagName).Each(func(i int, el *goquery.Selection) {
			section := strings.TrimPrefix(el.Text(), "\n")
			tabsStr := el.AttrOr(AttrTab, "0")
			tabs, err := strconv.Atoi(tabsStr)
			if err == nil && tabs > 0 {
				section = deTabulateSection(section, tabs)
			}
			fmt.Fprintln(out, section)
		})
	}
	return nil
}

func deTabulateSection(section string, tabs int) string {
	prefix := strings.Repeat("\t", tabs)
	var lines []string
	for _, line := range strings.Split(section, "\n") {
		lines = append(lines, strings.TrimPrefix(line, prefix))
	}
	return strings.Join(lines, "\n")
}
