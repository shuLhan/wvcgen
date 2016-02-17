// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"bytes"
	"compress/lzw"
	"fmt"
	"github.com/shuLhan/dsv"
)

/*
CompressRate is a feature that compute compression rate of inserted text.
*/
type CompressRate struct {
	dsv.Column
}

// init Register to list of feature
func init() {
	Register(&CompressRate{}, dsv.TReal, "compressrate")
}

/*
GetValues return feature values.
*/
func (ftr *CompressRate) GetValues() dsv.Column {
	return ftr.Column
}

/*
compressRateLzw return compression rate of `text` as

	length of compressed text / length of text
*/
func compressRateLzw(text string) (float64, error) {
	var buf bytes.Buffer
	btext := []byte(text)
	btextlen := len(btext)

	if btextlen <= 0 {
		return 0, nil
	}

	w := lzw.NewWriter(&buf, lzw.MSB, 8)

	n, e := w.Write(btext)
	if e != nil {
		fmt.Printf("error: %s, rate: %d, len: %d\n", e, n, btextlen)
		return 1, e
	}

	e = w.Close()
	if e != nil {
		fmt.Printf("error: %s, rate: %d, len: %d\n", e, n, btextlen)
		return 1, e
	}

	return float64(buf.Len()) / float64(btextlen), nil
}

/*
Compute compress rate of inserted text.
*/
func (ftr *CompressRate) Compute(dataset dsv.Dataset) {
	adds := dataset.GetColumnByName("additions")

	for _, rec := range adds.Records {
		r := &dsv.Record{}

		v, _ := compressRateLzw(rec.String())

		// round it to five digit after comma.
		r.SetFloat(float64(int(v*100000)) / 100000)

		ftr.PushBack(r)
	}
}
