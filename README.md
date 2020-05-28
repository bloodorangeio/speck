# Speck

[![GitHub Actions status](https://github.com/bloodorangeio/speck/workflows/build/badge.svg)](https://github.com/bloodorangeio/speck/actions?query=workflow%3Abuild) [![GoDoc](https://godoc.org/github.com/bloodorangeio/speck?status.svg)](https://godoc.org/github.com/bloodorangeio/speck)

![](https://raw.githubusercontent.com/bloodorangeio/speck/master/speck.png)

Speck is a tool that allows you to extract text between `<speck></speck>` tags found within one or more files.
It is agnostic to the format or extension of files used as input.
The result of the combined extracted text is printed to stdout.

This might be useful, for example, for building documentation from special comments placed in your source code.

## Installing

Requires Go 1.14+.

```
go build ./speck.go
sudo mv ./speck /usr/local/bin
```

## How to use

Start with a file containing `<speck></speck>` tags within, such as the following `main.go`:

```go
package main

/*
<speck>
# Hello World
</speck>
*/

import (
	"fmt"
)

func main() {
	/*
	<speck tab=1>
	This is a test.
	</speck>
	*/
	fmt.Println("hello world")
}
```

Then run `speck`, and direct the output to a file such as `example.md`:

```
speck main.go > example.md
```

The contents of `example.md`:
```
# Hello World

This is a test.

```

Notice the optional `tab` attribute, 
which will instruct Speck to trim x number of tabs from the left side of each line within a section.

Speck can also be used with multiple files, combining the output in order. 
You can also leverage shell globbing. Here is an example using some test files within this repo:

```
speck testdata/go-test-suite-example/*.go
```

which would be the same as running
```
speck testdata/go-test-suite-example/00_setup_test.go \
  testdata/go-test-suite-example/02_second_test.go \
  testdata/go-test-suite-example/01_first_test.go
```
