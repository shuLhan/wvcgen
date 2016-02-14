// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/dsv"
	"net"
)

// Anonim compute wether editor is login or from anonymous
// (logged by IP address).
type Anonim struct {
	dsv.Column
}

func init() {
	Register(&Anonim{}, dsv.TInteger, "anonim")
}

/*
GetValues return feature values.
*/
func (anon *Anonim) GetValues() dsv.Column {
	return anon.Column
}

/*
Compute if record in column is IP address then it is an anonim and set
their value to 1, otherwise set to 0.
*/
func (anon *Anonim) Compute(dataset dsv.Dataset) {
	col := dataset.GetColumnByName("editor")

	for _, rec := range col.Records {
		r := &dsv.Record{
			V: float64(0.0),
		}

		IP := net.ParseIP(rec.String())

		if IP != nil {
			r.SetFloat(1.0)
		}

		anon.PushBack(r)
	}
}
