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
PronounFrequency will count frequency of first and second person pronoun in
inserted text.
*/
type PronounFrequency Feature

func init() {
	Register(&PronounFrequency{}, tabula.TReal, "pronoun_frequency")
}

/*
Compute frequency of pronoun words in inserted text.
*/
func (ftr *PronounFrequency) Compute(dataset tabula.Dataset) {
	col := dataset.GetColumnByName("additions")

	for _, rec := range col.Records {
		text := rec.String()
		if len(text) == 0 {
			ftr.PushBack(&tabula.Record{V: float64(0)})
			continue
		}

		in := clean.WikiText(text)

		freq := tekstus.StringFrequenciesOf(in, tekstus.PronounWords,
			false)

		ftr.PushBack(&tabula.Record{V: Round(freq)})
	}
}
