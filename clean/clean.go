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

/*
WikiText will remove URI, wiki syntax and and mark-up, multiple spaces in
text and return it.
*/
func WikiText(text string) string {
	text = tekstus.StringRemoveURI(text)
	text = tekstus.StringRemoveWikiMarkup(text)
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
	text = strings.Replace(text, "<nowiki>", " ", -1)
	text = strings.Replace(text, "</nowiki>", " ", -1)
	text = strings.Replace(text, "&nbsp;", " ", -1)
	text = strings.TrimSpace(text)
	text = tekstus.StringMergeSpaces(text, true)

	return text
}

/*
WikiFile will remove URI, wiki syntax and and mark-up, multiple spaces in
content of file `in` and save it in file `out`.
*/
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
