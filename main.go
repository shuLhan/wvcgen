// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/tabula"
	"github.com/shuLhan/wvcgen/feature"
	wvcreader "github.com/shuLhan/wvcgen/reader"
	"github.com/shuLhan/wvcgen/revision"
	"io"
)

var (
	fInputDsv = "wvc2010_features.dsv"
)

/*
initReader set configuration of generator.
*/
func initReader(reader *wvcreader.Reader) {
	revision.SetDir(reader.RevisionDir)
	revision.SetCleanDir(reader.RevisionCleanDir)
}

/*
InitReadWriter initialize reader and writer.
*/
func InitReadWriter(finput string) (reader *wvcreader.Reader,
	writer *dsv.Writer) {
	reader, e := wvcreader.NewReader(finput)

	if e != nil {
		panic(e)
	}

	writer, e = dsv.NewWriter(finput)

	if e != nil {
		panic(e)
	}

	initReader(reader)

	return reader, writer
}

/*
computeFeatures will compute each feature listed in writer output metadata.
*/
func computeFeatures(reader *wvcreader.Reader, writer *dsv.Writer) (
	ftrValues *tabula.Dataset,
) {
	ftrValues = &tabula.Dataset{}

	ftrValues.Init(tabula.DatasetModeColumns, nil, nil)

	for _, md := range writer.OutputMetadata {
		fmt.Println(">>> computing feature", md.Name)

		ftr := feature.GetByName(md.Name)

		// No feature name found, search the column name in
		// input metadata.
		if ftr == nil {
			getAsInputColumn(reader, md.Name, ftrValues)
			continue
		}

		ftr.Compute(reader.Dataset)

		col := ftr.Interface().(*tabula.Column)

		ftrValues.PushColumn(*col)
	}

	return
}

/*
getAsInputColumn return feature values as in input column.
*/
func getAsInputColumn(reader *wvcreader.Reader, colName string,
	ftrValues *tabula.Dataset,
) {
	ftr := reader.GetColumnByName(colName)

	if ftr == nil {
		return
	}

	// Add column in input as feature
	ftrValues.PushColumn(*ftr)
}

/*
Generate start computing the feature values which has been defined in
input file `finput`.

If `featureName` is not empty, it will be added to list of feature that will
be computed.
*/
func Generate(featureName, finput string) {
	var ftrValues *tabula.Dataset
	var e error
	var n int

	reader, writer := InitReadWriter(finput)

	if featureName != "" {
		ftrMd := dsv.Metadata{
			Name: featureName,
		}

		writer.AddMetadata(ftrMd)
	}

	for {
		n, e = dsv.Read(reader)

		if n <= 0 {
			break
		}

		ftrValues = computeFeatures(reader, writer)

		_, ewrite := writer.WriteColumns(&ftrValues.Columns, nil)

		if ewrite != nil {
			panic(e)
		}

		if e == io.EOF {
			break
		}
	}

	e = reader.Close()
	if e != nil {
		panic(e)
	}

	e = writer.Close()

	if e != nil {
		panic(e)
	}
}

func main() {
	flag.Parse()

	if len(flag.Args()) >= 2 {
		fInputDsv = flag.Arg(2)
	}

	fmt.Println(">>> Processing", fInputDsv)

	Generate("", fInputDsv)
}
