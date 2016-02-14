// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/dsv"
)

/*
SizeRatio return the length of new revision substracted by length of old
revision.
*/
type SizeRatio struct {
	dsv.Column
}

func init() {
	// Register to list of feature
	Register(&SizeRatio{}, dsv.TInteger, "sizeratio")
}

/*
GetValues return feature values.
*/
func (ftr *SizeRatio) GetValues() dsv.Column {
	return ftr.Column
}

/*
Compute return the length of new revision substracted by length of old
revision.
*/
func (ftr *SizeRatio) Compute(dataset dsv.Dataset) {
	adds := dataset.GetColumnByName("additions")
	dels := dataset.GetColumnByName("deletions")

	delslen := dels.Len()

	for x, rec := range adds.Records {
		if x >= delslen {
			// Just in case additions is greater than deletions
			break
		}

		r := &dsv.Record{}

		newlen := len(rec.String())
		oldlen := len(dels.Records[x].String())
		ratio := float64(1+newlen) / float64(1+oldlen)

		r.SetFloat(ratio)

		ftr.PushBack(r)
	}
}
