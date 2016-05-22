// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"fmt"
	"github.com/shuLhan/tabula"
	"github.com/shuLhan/tekstus"
	"github.com/shuLhan/wvcgen/clean"
	"github.com/shuLhan/wvcgen/revision"
)

// TermFrequency compute frequency of words in inserted text againts the new
// revision.
type TermFrequency Feature

func init() {
	Register(&TermFrequency{}, tabula.TReal, "term_frequency")
}

/*
Compute the frequency of inserted words.
*/
func (ftr *TermFrequency) Compute(dataset tabula.DatasetInterface) {
	newrevidx := dataset.GetColumnByName("newrevisionid")
	adds := dataset.GetColumnByName("additions")
	recordslen := len(adds.Records)

	for x, rec := range adds.Records {
		r := tabula.NewRecordReal(float64(0))
		// Get inserted words.
		intext := rec.String()

		if len(intext) == 0 {
			ftr.PushBack(r)
			continue
		}

		intext = clean.WikiText(intext)
		inWords := tekstus.StringSplitWords(intext, true, true)

		// Get content of new revision.
		revid := newrevidx.Records[x].String()

		if DEBUG >= 2 {
			fmt.Printf("[feature] term_frequency: %d/%d processing %q\n",
				x, recordslen, revid)
		}

		newtext, e := revision.GetContentClean(revid)
		if e != nil {
			ftr.PushBack(r)
			continue
		}

		newWords := tekstus.StringSplitWords(newtext, true, false)

		freq := tekstus.WordsFrequenciesOf(newWords, inWords, false)

		r.SetFloat(Round(freq))

		ftr.PushBack(r)
	}
}
