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
	Register(&DigitRatio{}, dsv.TReal, "digitratio")
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
		r := &dsv.Record{}

		ratio := tekstus.RatioDigit(rec.String())

		// round it to five digit after comma.
		r.SetFloat(float64(int(ratio*100000)) / 100000)

		ftr.PushBack(r)
	}
}
