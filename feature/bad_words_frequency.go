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
	Register(&BadWordsFrequency{}, dsv.TInteger, "bad_words_frequency")
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
		in := clean.WikiText(rec.String())
		inWords := tekstus.StringSplitWords(in, true, false)

		freq := tekstus.WordsFrequenciesOf(inWords, tekstus.BadWords,
			false)

		freq = float64(int(freq*100000)) / 100000

		ftr.PushBack(&dsv.Record{V: freq})
	}
}
