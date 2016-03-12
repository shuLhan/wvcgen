// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/tabula"
)

/*
Class change the classification from text to numeric. The "regular" edit
will become 0 and the "vandalism" will become 1.
*/
type Class Feature

func init() {
	Register(&Class{}, tabula.TInteger, "class")
}

/*
Compute change the classification from text to numeric. The "regular" edit
will become 0 and the "vandalism" will become 1.
*/
func (ftr *Class) Compute(dataset tabula.DatasetInterface) {
	col := dataset.GetColumnByName("class")

	for _, rec := range col.Records {
		r := &tabula.Record{
			V: int64(0),
		}

		if rec.String() == "vandalism" {
			r.SetInteger(1)
		}

		ftr.PushBack(r)
	}
}
