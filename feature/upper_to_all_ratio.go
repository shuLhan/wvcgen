// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/tabula"
	"github.com/shuLhan/tekstus"
)

/*
UpperToAllRatio is a feature that compare uppercase with all characters.
*/
type UpperToAllRatio Feature

// init Register to list of feature
func init() {
	Register(&UpperToAllRatio{}, tabula.TReal, "upper_to_all_ratio")
}

/*
Compute ratio of uppercase to all characters in new revision.
*/
func (ftr *UpperToAllRatio) Compute(dataset tabula.DatasetInterface) {
	adds := dataset.GetColumnByName("additions")

	for _, rec := range adds.Records {
		ratio := tekstus.RatioUpper(rec.String())

		ftr.PushBack(tabula.NewRecordReal(Round(ratio)))
	}
}
