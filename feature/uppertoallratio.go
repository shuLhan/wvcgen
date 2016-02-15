// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/tekstus"
)

/*
UpperToAllRatio is a feature that compare uppercase with all characters.
*/
type UpperToAllRatio struct {
	dsv.Column
}

func init() {
	// Register to list of feature
	Register(&UpperToAllRatio{}, dsv.TReal, "uppertoallratio")
}

/*
GetValues return feature values.
*/
func (ftr *UpperToAllRatio) GetValues() dsv.Column {
	return ftr.Column
}

/*
Compute ratio of uppercase to all characters in new revision.
*/
func (ftr *UpperToAllRatio) Compute(dataset dsv.Dataset) {
	adds := dataset.GetColumnByName("additions")

	for _, rec := range adds.Records {
		r := &dsv.Record{}

		ratio := tekstus.RatioUpper(rec.String())

		// round it to five digit after comma.
		r.SetFloat(float64(int(ratio*100000)) / 100000)

		ftr.PushBack(r)
	}
}
