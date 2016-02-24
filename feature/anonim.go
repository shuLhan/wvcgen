// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/tabula"
	"net"
)

// Anonim compute wether editor is login or from anonymous
// (logged by IP address).
type Anonim Feature

func init() {
	Register(&Anonim{}, tabula.TInteger, "anonim")
}

/*
Compute if record in column is IP address then it is an anonim and set
their value to 1, otherwise set to 0.
*/
func (anon *Anonim) Compute(dataset tabula.Dataset) {
	col := dataset.GetColumnByName("editor")

	for _, rec := range col.Records {
		r := &tabula.Record{
			V: float64(0.0),
		}

		IP := net.ParseIP(rec.String())

		if IP != nil {
			r.SetFloat(1.0)
		}

		anon.PushBack(r)
	}
}
