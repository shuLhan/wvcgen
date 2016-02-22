// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/wvcgen/revision"
)

/*
SizeRatio is a feature that compare the size ratio of new with old revision.
*/
type SizeRatio struct {
	dsv.Column
}

func init() {
	// Register to list of feature
	Register(&SizeRatio{}, dsv.TReal, "size_ratio")
}

/*
GetValues return feature values.
*/
func (ftr *SizeRatio) GetValues() dsv.Column {
	return ftr.Column
}

/*
Compute ratio of size between new and old revision.
*/
func (ftr *SizeRatio) Compute(dataset dsv.Dataset) {
	oldid := dataset.GetColumnByName("oldrevisionid")
	newid := dataset.GetColumnByName("newrevisionid")

	oldidlen := newid.Len()

	for x, rec := range newid.Records {
		if x >= oldidlen {
			// Just in case additions is greater than deletions
			break
		}

		newlen := revision.GetSize(rec.String())
		oldlen := revision.GetSize(oldid.Records[x].String())
		difflen := float64(1+newlen) / float64(1+oldlen)

		ftr.PushBack(&dsv.Record{V: Round(difflen)})
	}
}
