// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/tabula"
	"github.com/shuLhan/wvcgen/revision"
	"math"
)

/*
SizeIncrement is a feature that compare the size of new with old revision by
subtracting their length.
*/
type SizeIncrement Feature

// init will register the featyre to the global list.
func init() {
	Register(&SizeIncrement{}, tabula.TInteger, "size_increment")
}

/*
Compute the absolute size increment.
*/
func (ftr *SizeIncrement) Compute(dataset tabula.Dataset) {
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
		difflen := math.Abs(float64(newlen - oldlen))

		ftr.PushBack(&tabula.Record{V: difflen})
	}
}
