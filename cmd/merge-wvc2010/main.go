// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/tabula"
	"github.com/shuLhan/tekstus/diff"
)

const (
	fEditsDsv           = "edits.dsv"
	fGoldAnnotationsDsv = "gold-annotations.dsv"
	fOutDsv             = "merge_edits_gold.dsv"
	dRevisions          = "../../pan-wvc-2010/revisions/"
)

/*
doDiff read old and new revisions from edit and compare both of them to get
deletions in old rev and additions in new rev.

Deletions and additions then combined into one string and appended to dataset.
*/
func doDiff(readset dsv.ReaderInterface) {
	diffset, e := dsv.NewReader("")

	if e != nil {
		panic(e)
	}

	e = diffset.SetDatasetMode(dsv.DatasetModeROWS)

	if e != nil {
		panic(e)
	}

	md := dsv.NewMetadata("deletions", "string", ",", "\"", "\"", nil)
	diffset.AddInputMetadata(md)

	md = dsv.NewMetadata("additions", "string", ",", "\"", "\"", nil)
	diffset.AddInputMetadata(md)

	for _, row := range readset.GetDataset().GetDataAsRows() {
		oldrevid := dRevisions + row[2].String() + ".txt"
		newrevid := dRevisions + row[3].String() + ".txt"

		diffs, e := diff.Files(oldrevid, newrevid, diff.LevelWords)

		if e != nil {
			panic(e)
		}

		dels := diffs.GetAllDels()
		delstr := dels.Join(" ")
		delrec, e := tabula.NewRecord(delstr, tabula.TString)

		if e != nil {
			panic(e)
		}

		adds := diffs.GetAllAdds()
		addstr := adds.Join(" ")
		addrec, e := tabula.NewRecord(addstr, tabula.TString)

		if e != nil {
			panic(e)
		}

		diffrow := tabula.Row{}

		diffrow.PushBack(delrec)
		diffrow.PushBack(addrec)

		diffset.PushRow(diffrow)
	}

	readset.MergeColumns(diffset)
}

func main() {
	readset, e := dsv.SimpleMerge(fEditsDsv, fGoldAnnotationsDsv)
	if e != nil {
		panic(e)
	}
	fmt.Printf(">>> merging %d rows\n", readset.GetDataset().GetNRow())

	fmt.Println(">>> diffing ...")
	doDiff(readset)

	fmt.Println(">>> writing ...")
	e = dsv.SimpleWrite(readset, fOutDsv)

	if e != nil {
		panic(e)
	}

	fmt.Printf(">>> writing %d rows\n", n)
err:
	_ = writer.Close()
}
