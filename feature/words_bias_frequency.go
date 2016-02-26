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
WordsBiasFrequency will count frequency of colloquial words with high bias in
inserted text.
*/
type WordsBiasFrequency Feature

func init() {
	Register(&WordsBiasFrequency{}, tabula.TReal, "words_bias_frequency")
}

/*
Compute frequency of biased words.
*/
func (ftr *WordsBiasFrequency) Compute(dataset tabula.Dataset) {
	col := dataset.GetColumnByName("additions")

	for _, rec := range col.Records {
		text := rec.String()
		if len(text) == 0 {
			ftr.PushBack(&tabula.Record{V: float64(0)})
			continue
		}

		in := clean.WikiText(text)

		freq := tekstus.StringFrequenciesOf(in,
			tekstus.BiasedWords, false)

		ftr.PushBack(&tabula.Record{V: Round(freq)})
	}
}
