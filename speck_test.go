package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	goTestSuiteExampleRootDir                = "./testdata/go-test-suite-example"
	goTestSuiteExampleExpectedResultFilename = "expected-result.md"

	htmlInHtmlFilename = "./testdata/html-in-html.txt"
	goodTabsFilename   = "./testdata/good-tabs.txt"
	badTabsFilename    = "./testdata/bad-tabs.txt"
)

func TestPrintFileSections(t *testing.T) {
	// Test bad file list
	out := new(bytes.Buffer)
	err := PrintFileSections([]string{"/some/nonexistent/path.txt"}, out)
	if err == nil {
		t.Fatal("expected error with bad filepath but got nil")
	}

	// Happy path
	expectedResultBytes, err := ioutil.ReadFile(
		filepath.Join(goTestSuiteExampleRootDir, goTestSuiteExampleExpectedResultFilename))
	if err != nil {
		t.Fatal(err)
	}
	expectedResult := string(expectedResultBytes)
	files, err := listFilesInDirWithExtension(goTestSuiteExampleRootDir, ".go")
	if err != nil {
		t.Fatal(err)
	}
	if len(files) == 0 {
		t.Fatal("unable to list .go files in test directory")
	}
	out = new(bytes.Buffer)
	err = PrintFileSections(files, out)
	if err != nil {
		t.Fatal(err)
	}
	result := out.String()
	if result != expectedResult {
		fmt.Println("RESULT:")
		fmt.Println(result)
		t.Fatal("Output from go-test-suite-example does not match expected output")
	}

	// Make sure HTML tags are not escaped during extraction
	out = new(bytes.Buffer)
	err = PrintFileSections([]string{htmlInHtmlFilename}, out)
	if err != nil {
		t.Fatal(err)
	}
	result = out.String()
	if result != "Example: <div>somestuff</div>\n\n" {
		fmt.Println("RESULT:")
		fmt.Println(result)
		t.Fatalf("Output from %s does not match expected output", htmlInHtmlFilename)
	}

	// Good tab get detabulated
	out = new(bytes.Buffer)
	err = PrintFileSections([]string{goodTabsFilename}, out)
	if err != nil {
		t.Fatal(err)
	}
	result = out.String()
	if result != "Good\n\n" {
		fmt.Println("RESULT:")
		fmt.Println(result)
		t.Fatalf("Output from %s does not match expected output", badTabsFilename)
	}

	// Bad tab throws error
	out = new(bytes.Buffer)
	err = PrintFileSections([]string{badTabsFilename}, out)
	if err == nil {
		t.Fatalf("Expected error parsing tabs in %s", badTabsFilename)
	}
}

func listFilesInDirWithExtension(dirname string, extension string) ([]string, error) {
	var files []string
	err := filepath.Walk(dirname, func(path string, stat os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !stat.IsDir() && strings.HasSuffix(path, extension) {
			files = append(files, path)
		}

		return nil
	})
	return files, err
}
