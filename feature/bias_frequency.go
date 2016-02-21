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
BiasFrequency will count frequency of colloquial words with high bias.
inserted text.
*/
type BiasFrequency struct {
	dsv.Column
}

func init() {
	Register(&BiasFrequency{}, dsv.TReal, "bias_frequency")
}

/*
GetValues return feature values.
*/
func (ftr *BiasFrequency) GetValues() dsv.Column {
	return ftr.Column
}

/*
Compute frequency of biased words.
*/
func (ftr *BiasFrequency) Compute(dataset dsv.Dataset) {
	col := dataset.GetColumnByName("additions")

	for _, rec := range col.Records {
		in := clean.WikiText(rec.String())
		inWords := tekstus.StringSplitWords(in, true, false)

		freq := tekstus.WordsFrequenciesOf(inWords,
			tekstus.BiasedWords, false)

		freq = float64(int(freq*100000)) / 100000

		ftr.PushBack(&dsv.Record{V: freq})
	}
}
