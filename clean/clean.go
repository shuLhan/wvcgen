// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package clean contain functions to clean wiki text, including removing wiki
syntax, links, and section numbering.
*/
package clean

import (
	"github.com/shuLhan/tekstus"
	"io/ioutil"
	"log"
	"strings"
)

func WikiText(text string) string {
	text = tekstus.StringRemoveURI(text)
	text = strings.Replace(text, "\r\n", "\n", -1)
	text = strings.Replace(text, "[", " ", -1)
	text = strings.Replace(text, "]", " ", -1)
	text = strings.Replace(text, "{", " ", -1)
	text = strings.Replace(text, "}", " ", -1)
	text = strings.Replace(text, "|", " ", -1)
	text = strings.Replace(text, "=", " ", -1)
	text = strings.Replace(text, "#", " ", -1)
	text = strings.Replace(text, "'s", " ", -1)
	text = strings.Replace(text, "'", " ", -1)
	text = strings.Replace(text, "<ref>", " ", -1)
	text = strings.Replace(text, "</ref>", " ", -1)
	text = strings.Replace(text, "<br />", " ", -1)
	text = strings.Replace(text, "<br/>", " ", -1)
	text = strings.Replace(text, "<br>", " ", -1)
	text = strings.TrimSpace(text)
	text = tekstus.StringMergeSpaces(text, true)

	return text
}

func WikiFile(in, out string) {
	bufIn, e := ioutil.ReadFile(in)
	if e != nil {
		log.Fatal(e)
	}

	newtext := WikiText(string(bufIn))

	e = ioutil.WriteFile(out, []byte(newtext), 0644)
	if e != nil {
		log.Fatal(e)
	}
}
