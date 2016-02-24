// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"fmt"
	"github.com/shuLhan/tabula"
	"github.com/shuLhan/tekstus"
	"github.com/shuLhan/wvcgen/clean"
)

// LongestWord find and return the longset word in inserted text.
type LongestWord Feature

func init() {
	Register(&LongestWord{}, tabula.TInteger, "longest_word")
}

/*
Compute the longest word in inserted text.
*/
func (ftr *LongestWord) Compute(dataset tabula.Dataset) {
	adds := dataset.GetColumnByName("additions")
	addslen := adds.Len()

	for x, rec := range adds.Records {
		text := rec.String()
		textlen := len(text)

		if textlen == 0 {
			ftr.PushBack(&tabula.Record{V: int64(0)})
			continue
		}

		text = clean.WikiText(text)
		inWords := tekstus.StringSplitWords(text, true, true)
		slong, _ := tekstus.WordsFindLongest(inWords)

		fmt.Printf(">>> %d/%d longest word: %q\n", x, addslen, slong)

		slonglen := int64(len(slong))

		ftr.PushBack(&tabula.Record{V: slonglen})
	}
}
