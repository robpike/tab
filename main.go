// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Tab applies the Elastic Tabs algorithm (see
// http://nickgravgaard.com/elastictabstops/index.html)nto the concatenation of its
// input files (default standard input), and writes the result to standard output
// using spaces to align the fields. In the input, columns must be separated by
// single tab characters. That is, given columnar input demarcated by single tabs,
// the tab program will align the columns for output, assuming a constant-width
// character set.
//
package main // import "robpike.io/cmd/tab"

import (
	"flag"
	"io"
	"log"
	"os"
	"text/tabwriter"
)

var (
	minWidth = flag.Int("wid", 4, "minimum `width` of an output column")
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("tab: ")
	flag.Parse()

	writer := tabwriter.NewWriter(os.Stdout, *minWidth, 4, 1, ' ', 0)
	if flag.NArg() == 0 {
		_, err := io.Copy(writer, os.Stdin)
		if err != nil {
			log.Fatalf("writing to tabwriter: %v", err)
		}
	} else {
		for _, name := range flag.Args() {
			fd, err := os.Open(name)
			if err != nil {
				log.Fatalf("%v", err)
			}
			_, err = io.Copy(writer, fd)
			if err != nil {
				log.Fatalf("writing %s to tabwriter: %v", name, err)
			}
			fd.Close()
		}
	}
	if err := writer.Flush(); err != nil {
		log.Fatalf("flushing tabwriter output: %v", err)
	}
}
