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
SexWordsFrequency will count frequency of non-vulgar, sex-related words.
*/
type SexWordsFrequency struct {
	dsv.Column
}

func init() {
	Register(&SexWordsFrequency{}, dsv.TReal, "sex_words_frequency")
}

/*
GetValues return feature values.
*/
func (ftr *SexWordsFrequency) GetValues() dsv.Column {
	return ftr.Column
}

/*
Compute frequency of sex related words.
*/
func (ftr *SexWordsFrequency) Compute(dataset dsv.Dataset) {
	col := dataset.GetColumnByName("additions")

	for _, rec := range col.Records {
		in := clean.WikiText(rec.String())

		if len(in) == 0 {
			ftr.PushBack(&dsv.Record{V: float64(0)})
			continue
		}

		inWords := tekstus.StringSplitWords(in, true, false)

		freq := tekstus.WordsFrequenciesOf(inWords,
			tekstus.SexWords, false)

		freq = float64(int(freq*100000)) / 100000

		ftr.PushBack(&dsv.Record{V: freq})
	}
}
