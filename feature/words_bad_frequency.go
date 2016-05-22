// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/tabula"
	"github.com/shuLhan/tekstus"
	"github.com/shuLhan/wvcgen/clean"
)

/*
WordsBadFrequency will compute frequency of bad words, colloquial words or
bad writing skill words.
*/
type WordsBadFrequency Feature

func init() {
	Register(&WordsBadFrequency{}, tabula.TReal, "words_bad_frequency")
}

/*
Compute frequency of bad words.
*/
func (ftr *WordsBadFrequency) Compute(dataset tabula.DatasetInterface) {
	col := dataset.GetColumnByName("additions")

	for _, rec := range col.Records {
		r := tabula.NewRecordReal(float64(0))
		intext := rec.String()

		if len(intext) == 0 {
			ftr.PushBack(r)
			continue
		}

		intext = clean.WikiText(intext)

		if len(intext) == 0 {
			ftr.PushBack(r)
			continue
		}

		inWords := tekstus.StringSplitWords(intext, true, false)

		freq := tekstus.WordsFrequenciesOf(inWords, tekstus.BadWords,
			false)

		r.SetFloat(Round(freq))

		ftr.PushBack(r)
	}
}
