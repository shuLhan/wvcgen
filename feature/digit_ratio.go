// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/tekstus"
)

/*
DigitRatio is a feature that compare digit to all charachter.
*/
type DigitRatio struct {
	dsv.Column
}

// init Register to list of feature
func init() {
	Register(&DigitRatio{}, dsv.TReal, "digit_ratio")
}

/*
GetValues return feature values.
*/
func (ftr *DigitRatio) GetValues() dsv.Column {
	return ftr.Column
}

/*
Compute calculate digit ratio in new revision.
*/
func (ftr *DigitRatio) Compute(dataset dsv.Dataset) {
	adds := dataset.GetColumnByName("additions")

	for _, rec := range adds.Records {
		ratio := tekstus.RatioDigit(rec.String())

		ftr.PushBack(&dsv.Record{V: Round(ratio)})
	}
}
