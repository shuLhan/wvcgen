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
PronounFrequency will count frequency of first and second person pronoun in
inserted text.
*/
type PronounFrequency struct {
	dsv.Column
}

func init() {
	Register(&PronounFrequency{}, dsv.TReal, "pronoun_frequency")
}

/*
GetValues return feature values.
*/
func (ftr *PronounFrequency) GetValues() dsv.Column {
	return ftr.Column
}

/*
Compute frequency of pronoun words in inserted text.
*/
func (ftr *PronounFrequency) Compute(dataset dsv.Dataset) {
	col := dataset.GetColumnByName("additions")

	for _, rec := range col.Records {
		text := rec.String()
		if len(text) == 0 {
			ftr.PushBack(&dsv.Record{V: float64(0)})
			continue
		}

		in := clean.WikiText(text)

		freq := tekstus.StringFrequenciesOf(in, tekstus.PronounWords,
			false)

		ftr.PushBack(&dsv.Record{V: Round(freq)})
	}
}
