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
VulgarFrequency will count frequency of vulgar words in inserted text.
*/
type VulgarFrequency Feature

func init() {
	Register(&VulgarFrequency{}, tabula.TReal, "vulgar_frequency")
}

/*
Compute frequency vulgar words in inserted text.
*/
func (ftr *VulgarFrequency) Compute(dataset tabula.Dataset) {
	col := dataset.GetColumnByName("additions")

	for _, rec := range col.Records {
		s := clean.WikiText(rec.String())

		freq := tekstus.StringFrequenciesOf(s, tekstus.VulgarWords,
			false)

		ftr.PushBack(&tabula.Record{V: Round(freq)})
	}
}
