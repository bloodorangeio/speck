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

	htmlInHtmlFilename       = "./testdata/html-in-html.txt"
	goodTabsFilename         = "./testdata/good-tabs.txt"
	badTabsFilename          = "./testdata/bad-tabs.txt"
	unclosedHTMLTagsFilename = "./testdata/unclosed-html-tags.txt"
	disappearingTagFilename = "./testdata/disappearing-tag.txt"
)

func TestBadFileList(t *testing.T) {
	// Test bad file list
	out := new(bytes.Buffer)
	err := PrintFileSections([]string{"/some/nonexistent/path.txt"}, out)
	if err == nil {
		t.Fatal("expected error with bad filepath but got nil")
	}
}

func TestHappyPath(t *testing.T) {
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
	out := new(bytes.Buffer)
	err = PrintFileSections(files, out)
	if err != nil {
		t.Fatal(err)
	}
	result := out.String()
	if result != expectedResult {
		fmt.Println("RESULT:")
		fmt.Println(result)
		t.Error("Output from go-test-suite-example does not match expected output")
	}
}

func TestHtmlInHtml(t *testing.T) {
	// Make sure HTML tags are not escaped during extraction
	out := new(bytes.Buffer)
	err := PrintFileSections([]string{htmlInHtmlFilename}, out)
	if err != nil {
		t.Fatal(err)
	}
	result := out.String()
	if result != "Example: <div>somestuff</div>\n\n" {
		fmt.Println("RESULT:")
		fmt.Println(result)
		t.Errorf("Output from %s does not match expected output", htmlInHtmlFilename)
	}
}

func TestGoodTabAttribute(t *testing.T) {
	// Good tab attribute gets detabulated
	out := new(bytes.Buffer)
	err := PrintFileSections([]string{goodTabsFilename}, out)
	if err != nil {
		t.Fatal(err)
	}
	result := out.String()
	if result != "Good\n\n" {
		fmt.Println("RESULT:")
		fmt.Println(result)
		t.Errorf("Output from %s does not match expected output", badTabsFilename)
	}
}

func TestBadTabAttribute(t *testing.T) {
	// Bad tab throws error
	out := new(bytes.Buffer)
	err := PrintFileSections([]string{badTabsFilename}, out)
	if err == nil {
		t.Fatalf("Expected error parsing tabs in %s", badTabsFilename)
	}
}

func TestUnclosedHtmlTags(t *testing.T) {
	// Make sure HTML tags within a section are treated as plain text
	out := new(bytes.Buffer)
	err := PrintFileSections([]string{unclosedHTMLTagsFilename}, out)
	if err != nil {
		t.Fatal(err)
	}
	result := out.String()
	if result != `Here is an example HTTP request:
`+"```"+`
GET /v2/<name>/blobs/<digest>
Host: <registry host>
Authorization: <scheme> <token>
`+"```\n\n" {
		fmt.Println("RESULT:")
		fmt.Println(result)
		t.Errorf("Output from %s does not match expected output", unclosedHTMLTagsFilename)
	}
}

func TestDisappearingTag(t *testing.T) {
	// Demonstrate inner HTML caveat
	out := new(bytes.Buffer)
	err := PrintFileSections([]string{disappearingTagFilename}, out)
	if err != nil {
		t.Fatal(err)
	}
	result := out.String()
	if result != "This is my API request: `POST /api/stuff/`\n" {
		fmt.Println("RESULT:")
		fmt.Println(result)
		t.Errorf("Output from %s does not match expected output", disappearingTagFilename)
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
