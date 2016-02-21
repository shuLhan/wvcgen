// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/wvcgen/feature"
	"github.com/shuLhan/wvcgen/reader"
	"github.com/shuLhan/wvcgen/revision"
	"io"
)

const (
	fInputDsv = "features.dsv"
)

/*
Init set configuration of generator.
*/
func Init(reader *reader.Reader) {
	revision.SetDir(reader.RevisionDir)
	revision.SetCleanDir(reader.RevisionCleanDir)
}

func main() {
	ftrValues := dsv.Dataset{}

	ftrValues.Init(dsv.DatasetModeColumns, nil, nil)

	reader, e := reader.NewReader(fInputDsv)

	if e != nil {
		panic(e)
		return
	}

	writer, e := dsv.NewWriter(fInputDsv)

	if e != nil {
		panic(e)
		return
	}

	Init(reader)

	for {
		n, e := dsv.Read(reader)

		if n <= 0 {
			break
		}

		computeFeatures(reader, writer, &ftrValues)

		if e == io.EOF {
			break
		}
	}

	e = reader.Close()
	if e != nil {
		panic(e)
	}

	_, e = writer.WriteColumns(&ftrValues.Columns, nil)

	if e != nil {
		panic(e)
	}

	e = writer.Close()

	if e != nil {
		panic(e)
	}
}

func computeFeatures(reader *reader.Reader, writer *dsv.Writer, ftrValues *dsv.Dataset) {
	for _, md := range writer.OutputMetadata {
		ftr := feature.ListFeatureGetByName(md.Name)

		// No feature name found, search the column name in
		// input metadata.
		if ftr == nil {
			getAsInputColumn(reader, md.Name, ftrValues)
			continue
		}

		fmt.Println(">>> computing feature", ftr.GetName())
		ftr.Compute(reader.Dataset)

		ftrValues.PushColumn(ftr.GetValues())
	}
}

func getAsInputColumn(reader *reader.Reader, colName string,
	ftrValues *dsv.Dataset,
) {
	ftr := reader.GetColumnByName(colName)

	if ftr == nil {
		return
	}

	// Add column in input as feature
	ftrValues.PushColumn(*ftr)
}
