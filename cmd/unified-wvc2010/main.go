// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/wvcgen/revision"
	"log"
	"time"
)

const (
	fEditsDsv           = "edits.dsv"
	fGoldAnnotationsDsv = "gold-annotations.dsv"
	fOutDsv             = "unified-wvc2010.dsv"
	dRevisions          = "../../pan-wvc-2010/revisions/"
)

func trace(s string) (string, time.Time) {
	log.Println("START:", s)
	return s, time.Now()
}

func un(s string, startTime time.Time) {
	endTime := time.Now()
	log.Println("  END: ", endTime.Sub(startTime))
}

/*
doDiff read old and new revisions from edit and compare both of them to get
deletions in old rev and additions in new rev.

Deletions and additions then combined into one string and appended to dataset.
*/
func doDiff(readset dsv.ReaderInterface) {
	ds := readset.GetDataset()
	ds.TransposeToColumns()

	oldids := ds.GetColumnByName("oldrevisionid").ToStringSlice()
	newids := ds.GetColumnByName("newrevisionid").ToStringSlice()

	revision.SetDir(dRevisions)

	diffset, e := revision.Diff(oldids, newids, ".txt")
	if e != nil {
		panic(e)
	}

	// Create input metadata for diff
	md := dsv.NewMetadata("deletions", "string", ",", "\"", "\"", nil)
	readset.AddInputMetadata(md)

	md = dsv.NewMetadata("additions", "string", ",", "\"", "\"", nil)
	readset.AddInputMetadata(md)

	ds.MergeColumns(diffset)
}

func main() {
	defer un(trace("Unified PAN-WVC-2010"))

	readset, e := dsv.SimpleMerge(fEditsDsv, fGoldAnnotationsDsv)
	if e != nil {
		panic(e)
	}
	fmt.Printf(">>> merging %d rows\n", readset.GetDataset().GetNRow())

	fmt.Println(">>> diffing ...")
	doDiff(readset)

	fmt.Println(">>> writing ...")
	n, e := dsv.SimpleWrite(readset, fOutDsv)
	if e != nil {
		panic(e)
	}
	fmt.Printf(">>> writing %d rows\n", n)
}
