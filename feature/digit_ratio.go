// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/tabula"
	"github.com/shuLhan/tekstus"
)

/*
DigitRatio is a feature that compare digit to all character.
*/
type DigitRatio Feature

// init Register to list of feature
func init() {
	Register(&DigitRatio{}, tabula.TReal, "digit_ratio")
}

/*
Compute calculate digit ratio in new revision.
*/
func (ftr *DigitRatio) Compute(dataset tabula.DatasetInterface) {
	adds := dataset.GetColumnByName("additions")

	for _, rec := range adds.Records {
		ratio := tekstus.RatioDigit(rec.String())

		ftr.PushBack(tabula.NewRecordReal(Round(ratio)))
	}
}
