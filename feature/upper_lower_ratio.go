// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/tabula"
	"github.com/shuLhan/tekstus"
)

/*
UpperLowerRatio is a feature that compare uppercase and lowercase characters.
*/
type UpperLowerRatio Feature

// init Register to list of feature
func init() {
	Register(&UpperLowerRatio{}, tabula.TReal, "upper_lower_ratio")
}

/*
Compute ratio of uppercase and lowercase in new revision.
*/
func (ftr *UpperLowerRatio) Compute(dataset tabula.DatasetInterface) {
	adds := dataset.GetColumnByName("additions")

	for _, rec := range adds.Records {
		ratio := tekstus.RatioUpperLowerChar(rec.String())

		ftr.PushBack(&tabula.Record{V: Round(ratio)})
	}
}
