// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"fmt"
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/tekstus"
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

	for x, rec := range adds.Records {
		r := &dsv.Record{}

		// Get inserted text.
		inText := rec.String()

		// Get content of new revision
		revid := newrevidx.Records[x].String()
		fmt.Printf(">>> %d processing new revision id %q\n", x, revid)

		newText, e := revision.GetContent(revid)
		if e != nil {
			r.SetFloat(0)
			ftr.PushBack(r)
			continue
		}

		newWords := tekstus.StringSplitWords(newText, true, false)
		inWords := tekstus.StringSplitWords(inText, true, true)

		freq := tekstus.ListStringFrequency(newWords, inWords, false)

		// round it to five digit after comma.
		r.SetFloat(float64(int(freq*100000)) / 100000)

		ftr.PushBack(r)
	}
}
