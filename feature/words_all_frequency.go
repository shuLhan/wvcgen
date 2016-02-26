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
WordsAllFrequency compute vandalism, pronouns, bias, sex, and bad words in
inserted text.
*/
type WordsAllFrequency Feature

func init() {
	Register(&WordsAllFrequency{}, tabula.TReal, "words_all_frequency")
}

/*
Compute frequency of all words.
*/
func (ftr *WordsAllFrequency) Compute(dataset tabula.Dataset) {
	allWords := GetAllWordList()

	col := dataset.GetColumnByName("additions")

	for _, rec := range col.Records {
		r := tabula.Record{V: float64(0)}

		s := rec.String()
		if len(s) == 0 {
			ftr.PushBack(&r)
			continue
		}

		s = clean.WikiText(s)
		if len(s) == 0 {
			ftr.PushBack(&r)
			continue
		}

		inWords := tekstus.StringSplitWords(s, true, false)

		freq := tekstus.WordsFrequenciesOf(inWords, allWords, false)

		r.SetFloat(Round(freq))

		ftr.PushBack(&r)
	}
}
