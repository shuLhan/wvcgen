// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/dsv"
)

/*
Class change the classification from text to numeric. The "regular" edit
will become 0 and the "vandalism" will become 1.
*/
type Class struct {
	dsv.Column
}

func init() {
	ftr := Class{
		Column: dsv.Column{
			Type: dsv.TInteger,
			Name: "class",
		},
	}

	// Register to list of feature
	ListFeatureAdd(&ftr)
}

/*
GetValues return feature values.
*/
func (ftr *Class) GetValues() dsv.Column {
	return ftr.Column
}

/*
Compute change the classification from text to numeric. The "regular" edit
will become 0 and the "vandalism" will become 1.
*/
func (ftr *Class) Compute(dataset dsv.Dataset) {
	col := dataset.GetColumnByName("class")

	for _, rec := range col.Records {
		r := &dsv.Record{
			V: int64(0),
		}

		if rec.String() == "vandalism" {
			r.SetInteger(1)
		}

		ftr.PushBack(r)
	}
}
