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
VulgarFrequency will count frequency of vulgar words in inserted text.
*/
type VulgarFrequency struct {
	dsv.Column
}

func init() {
	Register(&VulgarFrequency{}, dsv.TInteger, "vulgar_frequency")
}

/*
GetValues return feature values.
*/
func (ftr *VulgarFrequency) GetValues() dsv.Column {
	return ftr.Column
}

/*
Compute frequency vulgar words in inserted text.
*/
func (ftr *VulgarFrequency) Compute(dataset dsv.Dataset) {
	col := dataset.GetColumnByName("additions")

	for _, rec := range col.Records {
		s := clean.WikiText(rec.String())

		freq := tekstus.StringFrequenciesOf(s, tekstus.VulgarWords,
			false)

		freq = float64(int(freq*100000)) / 100000

		ftr.PushBack(&dsv.Record{V: freq})
	}
}
