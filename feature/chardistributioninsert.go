// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/wvcgen/revision"
)

/*
CharDistributionInsert is a feature that measure divergence of the character
distribution of the inserted text with respect to the expectation.
*/
type CharDistributionInsert struct {
	dsv.Column
}

// init Register to list of feature
func init() {
	Register(&CharDistributionInsert{}, dsv.TReal,
		"chardistributioninsert")
}

/*
GetValues return feature values.
*/
func (ftr *CharDistributionInsert) GetValues() dsv.Column {
	return ftr.Column
}

/*
Compute character distribution of inserted text.
*/
func (ftr *CharDistributionInsert) Compute(dataset dsv.Dataset) {
	oldrevid := dataset.GetColumnByName("oldrevisionid")
	adds := dataset.GetColumnByName("additions")

	for x, rold := range oldrevid.Records {
		r := &dsv.Record{}

		// count distribution of old revision
		oldText, e := revision.GetContent(rold.String())
		if e != nil {
			continue
		}

		// count distribution of inserted text
		inText := adds.Records[x].String()

		divergence := KullbackLeiblerDivergence(oldText, inText)

		// round it to five digit after comma.
		r.SetFloat(float64(int(divergence*100000)) / 100000)

		ftr.PushBack(r)
	}
}
