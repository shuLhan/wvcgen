// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"fmt"
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/tekstus"
	"github.com/shuLhan/wvcgen/clean"
	"github.com/shuLhan/wvcgen/revision"
)

// TermFrequency compute frequency of words in inserted text againts the new
// revision.
type TermFrequency struct {
	dsv.Column
}

func init() {
	Register(&TermFrequency{}, dsv.TReal, "term_frequency")
}

/*
GetValues return feature values.
*/
func (ftr *TermFrequency) GetValues() dsv.Column {
	return ftr.Column
}

/*
Compute the frequency of inserted words.
*/
func (ftr *TermFrequency) Compute(dataset dsv.Dataset) {
	newrevidx := dataset.GetColumnByName("newrevisionid")
	adds := dataset.GetColumnByName("additions")
	recordslen := len(adds.Records)

	for x, rec := range adds.Records {
		// Get inserted words.
		intext := rec.String()

		if len(intext) == 0 {
			ftr.PushBack(&dsv.Record{V: float64(0)})
			continue
		}

		intext = clean.WikiText(intext)
		inWords := tekstus.StringSplitWords(intext, true, true)

		// Get content of new revision.
		revid := newrevidx.Records[x].String()
		fmt.Printf(">>> term_frequency: %d/%d processing %q\n", x,
			recordslen, revid)

		newtext, e := revision.GetContentClean(revid)
		if e != nil {
			ftr.PushBack(&dsv.Record{V: float64(0)})
			continue
		}

		newWords := tekstus.StringSplitWords(newtext, true, false)

		freq := tekstus.WordsFrequenciesOf(newWords, inWords, false)

		ftr.PushBack(&dsv.Record{V: Round(freq)})
	}
}
