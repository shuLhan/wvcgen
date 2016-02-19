// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/shuLhan/wvcgen/clean"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("%s revisions-directory clean-directory\n", os.Args[0])
		return
	}

	revDir := os.Args[1]
	cleanDir := os.Args[2]

	files, e := ioutil.ReadDir(revDir)
	if e != nil {
		log.Fatal(e)
	}

	fileslen := len(files)

	for x, f := range files {
		if f.IsDir() {
			continue
		}

		in := revDir + "/" + f.Name()
		out := cleanDir + "/" + f.Name()

		log.Printf(">>> cleaning %d/%d: %s", x, fileslen, in)
		clean.WikiFile(in, out)
	}
}
