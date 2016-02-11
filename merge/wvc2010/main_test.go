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

const (
	fMainTestDsv = "main_test.dsv"
)

func TestMainOutput(t *testing.T) {
	rw, e := dsv.New(fMainTestDsv)

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
