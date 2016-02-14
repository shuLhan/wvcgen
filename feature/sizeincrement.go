// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/dsv"
)

/*
SizeIncrement is a feature that compare the size of new with old revision.
*/
type SizeIncrement struct {
	dsv.Column
}

func init() {
	// Register to list of feature
	Register(&SizeIncrement{}, dsv.TInteger, "sizeincrement")
}

/*
GetValues return feature values.
*/
func (ftr *SizeIncrement) GetValues() dsv.Column {
	return ftr.Column
}

/*
Compute the size increment by substracting length of addition with the length
of deletion.
*/
func (ftr *SizeIncrement) Compute(dataset dsv.Dataset) {
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

		r.SetInteger(int64(newlen - oldlen))

		ftr.PushBack(r)
	}
}
