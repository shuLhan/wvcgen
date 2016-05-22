// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"bytes"
	"compress/lzw"
	"fmt"
	"github.com/shuLhan/tabula"
)

/*
CompressRate is a feature that compute compression rate of inserted text.
*/
type CompressRate Feature

// init Register to list of feature
func init() {
	Register(&CompressRate{}, tabula.TReal, "compress_rate")
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
func (ftr *CompressRate) Compute(dataset tabula.DatasetInterface) {
	adds := dataset.GetColumnByName("additions")

	for _, rec := range adds.Records {
		v, _ := compressRateLzw(rec.String())

		ftr.PushBack(tabula.NewRecordReal(Round(v)))
	}
}
