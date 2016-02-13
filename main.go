// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
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

		computeFeatures(dsvRW, &ftrValues)

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

func computeFeatures(dsvRW *dsv.ReadWriter, ftrValues *dsv.Dataset) {
	for _, md := range dsvRW.OutputMetadata {
		ftr := feature.ListFeatureGetByName(md.Name)

		// No feature name found, search the column name in
		// input metadata.
		if ftr == nil {
			getAsInputColumn(dsvRW, md.Name, ftrValues)
			continue
		}

		fmt.Println(">>> computing feature", ftr.GetName())
		ftr.Compute(dsvRW.Dataset)

		ftrValues.PushColumn(ftr.GetValues())
	}
}

func getAsInputColumn(dsvRW *dsv.ReadWriter, colName string,
	ftrValues *dsv.Dataset,
) {
	ftr := dsvRW.GetColumnByName(colName)

	if ftr == nil {
		return
	}

	// Add column in input as feature
	ftrValues.PushColumn(*ftr)
}
