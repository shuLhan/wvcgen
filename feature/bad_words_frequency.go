// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/tekstus"
	"github.com/shuLhan/wvcgen/clean"
)

/*
BadWordsFrequency will compute frequency of bad words, colloquial words or
bad writing skill words.
*/
type BadWordsFrequency struct {
	dsv.Column
}

func init() {
	Register(&BadWordsFrequency{}, dsv.TReal, "bad_words_frequency")
}

/*
GetValues return feature values.
*/
func (ftr *BadWordsFrequency) GetValues() dsv.Column {
	return ftr.Column
}

/*
Compute frequency of bad words.
*/
func (ftr *BadWordsFrequency) Compute(dataset dsv.Dataset) {
	col := dataset.GetColumnByName("additions")

	for _, rec := range col.Records {
		intext := rec.String()

		if len(intext) == 0 {
			ftr.PushBack(&dsv.Record{V: float64(0)})
			continue
		}

		intext = clean.WikiText(intext)

		if len(intext) == 0 {
			ftr.PushBack(&dsv.Record{V: float64(0)})
			continue
		}

		inWords := tekstus.StringSplitWords(intext, true, false)

		freq := tekstus.WordsFrequenciesOf(inWords, tekstus.BadWords,
			false)

		ftr.PushBack(&dsv.Record{V: Round(freq)})
	}
}
