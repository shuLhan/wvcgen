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
BadWordsFrequency will compute frequency of bad words, colloquial words or
bad writing skill words.
*/
type BadWordsFrequency Feature

func init() {
	Register(&BadWordsFrequency{}, tabula.TReal, "bad_words_frequency")
}

/*
Compute frequency of bad words.
*/
func (ftr *BadWordsFrequency) Compute(dataset tabula.Dataset) {
	col := dataset.GetColumnByName("additions")

	for _, rec := range col.Records {
		intext := rec.String()

		if len(intext) == 0 {
			ftr.PushBack(&tabula.Record{V: float64(0)})
			continue
		}

		intext = clean.WikiText(intext)

		if len(intext) == 0 {
			ftr.PushBack(&tabula.Record{V: float64(0)})
			continue
		}

		inWords := tekstus.StringSplitWords(intext, true, false)

		freq := tekstus.WordsFrequenciesOf(inWords, tekstus.BadWords,
			false)

		ftr.PushBack(&tabula.Record{V: Round(freq)})
	}
}
