// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main_test

import (
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/dsv/util/assert"
	"io"
	"testing"
)

func TestMainOutput(t *testing.T) {
	rw, e := dsv.New("pan-wvc-2010_merge_edits_gold_test.dsv")

	if e != nil {
		t.Fatal(e)
	}

	for {
		n, e := dsv.Read(&rw.Reader)

		if e == io.EOF {
			_, e = rw.Write(&rw.Reader)
			break
		}
		if n > 0 {
			_, e = rw.Write(&rw.Reader)

			if e != nil {
				t.Fatal(e)
			}
		}
	}

	e = rw.Writer.Flush()
	if e != nil {
		t.Fatal(e)
	}

	assert.EqualFileContent(t, rw.GetInput(), rw.GetOutput())
}
