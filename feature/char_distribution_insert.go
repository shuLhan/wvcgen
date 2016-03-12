// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/tabula"
	"github.com/shuLhan/wvcgen/revision"
)

/*
CharDistributionInsert measure divergence of the character
distribution of the inserted text with respect to the expectation.
*/
type CharDistributionInsert Feature

// init Register to list of feature
func init() {
	Register(&CharDistributionInsert{}, tabula.TReal,
		"char_distribution_insert")
}

/*
Compute character distribution of inserted text.
*/
func (ftr *CharDistributionInsert) Compute(dataset tabula.DatasetInterface) {
	oldrevid := dataset.GetColumnByName("oldrevisionid")
	adds := dataset.GetColumnByName("additions")

	for x, rold := range oldrevid.Records {
		// count distribution of old revision
		oldText, e := revision.GetContent(rold.String())

		if e != nil {
			ftr.PushBack(&tabula.Record{V: 0.0})
			continue
		}

		// count distribution of inserted text
		inText := adds.Records[x].String()

		divergence := KullbackLeiblerDivergence(oldText, inText)

		ftr.PushBack(&tabula.Record{V: Round(divergence)})
	}
}
