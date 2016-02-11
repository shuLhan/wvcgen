// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/wvcgen/feature"
	"io"
)

const (
	fInputDsv = "features.dsv"
)

func main() {
	ftrValues := dsv.Dataset{}

	ftrValues.Init(dsv.DatasetModeColumns, nil, nil)

	dsvRW, e := dsv.New(fInputDsv)

	if e != nil {
		panic(e)
		return
	}

	for {
		n, e := dsv.Read(&dsvRW.Reader)

		if n <= 0 {
			break
		}

		for _, md := range dsvRW.OutputMetadata {
			ftr := feature.ListFeatureGetByName(md.Name)

			ftr.Compute(dsvRW.Dataset)

			ftrValues.PushColumn(ftr.GetValues())
		}

		if e == io.EOF {
			break
		}
	}

	_, e = dsvRW.WriteColumns(&ftrValues.Columns, nil)

	if e != nil {
		panic(e)
	}

	e = dsvRW.Close()

	if e != nil {
		panic(e)
	}
}
