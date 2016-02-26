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
WordsSexFrequency will count frequency of non-vulgar, sex-related words.
*/
type WordsSexFrequency Feature

func init() {
	Register(&WordsSexFrequency{}, tabula.TReal, "words_sex_frequency")
}

/*
Compute frequency of sex related words.
*/
func (ftr *WordsSexFrequency) Compute(dataset tabula.Dataset) {
	col := dataset.GetColumnByName("additions")

	for _, rec := range col.Records {
		text := rec.String()

		if len(text) == 0 {
			ftr.PushBack(&tabula.Record{V: float64(0)})
			continue
		}

		in := clean.WikiText(text)

		if len(in) == 0 {
			ftr.PushBack(&tabula.Record{V: float64(0)})
			continue
		}

		inWords := tekstus.StringSplitWords(in, true, false)

		freq := tekstus.WordsFrequenciesOf(inWords,
			tekstus.SexWords, false)

		ftr.PushBack(&tabula.Record{V: Round(freq)})
	}
}
