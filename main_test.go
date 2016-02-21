// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main_test

import (
	"fmt"
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/wvcgen/feature"
	wvcreader "github.com/shuLhan/wvcgen/reader"
	"github.com/shuLhan/wvcgen/revision"
	"io"
	"testing"
)

const (
	fInputDsv = "features_test.dsv"
)

/*
initReader set configuration of generator.
*/
func initReader(reader *wvcreader.Reader) {
	revision.SetDir(reader.RevisionDir)
	revision.SetCleanDir(reader.RevisionCleanDir)
}

/*
initReadWriter initialize reader and writer.
*/
func initReadWriter() (reader *wvcreader.Reader,
	writer *dsv.Writer) {
	reader, e := wvcreader.NewReader(fInputDsv)

	if e != nil {
		panic(e)
	}

	writer, e = dsv.NewWriter(fInputDsv)

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
	ftrValues *dsv.Dataset,
) {
	ftrValues = &dsv.Dataset{}

	ftrValues.Init(dsv.DatasetModeColumns, nil, nil)

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

	return
}

/*
getAsInputColumn return feature values as in input column.
*/
func getAsInputColumn(reader *wvcreader.Reader, colName string,
	ftrValues *dsv.Dataset,
) {
	ftr := reader.GetColumnByName(colName)

	if ftr == nil {
		return
	}

	// Add column in input as feature
	ftrValues.PushColumn(*ftr)
}

func callMain(t *testing.T, featureName string) {
	var ftrValues *dsv.Dataset
	var e error
	var n int

	reader, writer := initReadWriter()

	ftrMd := dsv.Metadata{
		Name: featureName,
	}

	writer.AddMetadata(ftrMd)

	for {
		n, e = dsv.Read(reader)

		if n <= 0 {
			break
		}

		ftrValues = computeFeatures(reader, writer)

		if e == io.EOF {
			break
		}
	}

	e = reader.Close()
	if e != nil {
		t.Fatal(e)
	}

	_, e = writer.WriteColumns(&ftrValues.Columns, nil)

	if e != nil {
		t.Fatal(e)
	}

	e = writer.Close()

	if e != nil {
		t.Fatal(e)
	}
}

func TestBiasFrequency(t *testing.T) {
	callMain(t, "bias_frequency")
}

func TestSexWordsFrequency(t *testing.T) {
	callMain(t, "sex_words_frequency")
}

func TestBadWordsFrequency(t *testing.T) {
	callMain(t, "bad_words_frequency")
}
