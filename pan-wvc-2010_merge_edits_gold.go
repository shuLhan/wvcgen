// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/tekstus/diff"
	"io"
)

/*
mergeDatasets merge edits.csv and gold-annotations.csv.

edits.csv contains edit metadata.
gold-annotations.csv contains classification (regular or vandalism) of edit.
*/
func mergeDatasets() (readset *dsv.Reader, e error) {
	var readgold *dsv.Reader

	// read edits
	fmt.Println(">>> read edits.csv")
	readedits, e := dsv.NewReader("pan-wvc-2010_edits.dsv")

	if e != nil {
		return
	}

	n, e := dsv.Read(readedits)

	if e != nil && e != io.EOF {
		goto End
	}

	fmt.Printf(">>> read %d rows\n", n)

	// read classifications
	fmt.Println(">>> read gold-annotations.csv")
	readgold, e = dsv.NewReader("pan-wvc-2010_gold-annotations.dsv")

	if e != nil {
		goto End
	}

	n, e = dsv.Read(readgold)

	if e != nil && e != io.EOF {
		goto End
	}

	fmt.Printf(">>> read %d rows\n", n)

	// Merge edits and gold annotation to get the class
	readedits.MergeColumns(readgold)
End:
	readedits.Close()
	readgold.Close()

	if e == io.EOF {
		e = nil
	}

	return readedits, e
}

/*
doDiff read old and new revisions from edit and compare both of them to get
deletions in old rev and additions in new rev.

Deletions and additions then combined into one string and appended to dataset.
*/
func doDiff(readset *dsv.Reader) {
	diffset, e := dsv.NewReader("")

	if e != nil {
		panic(e)
	}

	diffset.SetDatasetMode(dsv.DatasetModeROWS)

	md := dsv.NewMetadata("deletions", "string", ",", "\"", "\"", nil)
	diffset.AddInputMetadata(md)

	md = dsv.NewMetadata("additions", "string", ",", "\"", "\"", nil)
	diffset.AddInputMetadata(md)

	for _, row := range readset.GetDataAsRows() {
		oldrevid := "pan-wvc-2010/revisions/" + row[2].String() + ".txt"
		newrevid := "pan-wvc-2010/revisions/" + row[3].String() + ".txt"

		diffs, e := diff.Files(oldrevid, newrevid, diff.LevelWords)

		if e != nil {
			panic(e)
		}

		dels := diffs.GetAllDels()
		delstr := dels.Join(" ")
		delrec, e := dsv.NewRecord(delstr, dsv.TString)

		if e != nil {
			panic(e)
		}

		adds := diffs.GetAllAdds()
		addstr := adds.Join(" ")
		addrec, e := dsv.NewRecord(addstr, dsv.TString)

		if e != nil {
			panic(e)
		}

		diffrow := dsv.Row{}

		diffrow.PushBack(delrec)
		diffrow.PushBack(addrec)

		diffset.PushRow(diffrow)
	}

	readset.MergeColumns(diffset)
}

func main() {
	readset, e := mergeDatasets()

	if e != nil {
		fmt.Println("mergeDataset: ", e)
		panic(e)
	}

	fmt.Println(">>> diffing ...")
	doDiff(readset)

	writer, e := dsv.NewWriter("pan-wvc-2010_merge_edits_gold.dsv")

	if e != nil {
		fmt.Println("dsv.NewWriter: ", e)
		panic(e)
	}

	fmt.Println(">>> writing ...")
	n, e := writer.Write(readset)

	if e != nil {
		fmt.Println("writer.Write: ", e)
		goto End
	}

	fmt.Printf(">>> writing %d rows\n", n)
End:
	writer.Close()
}
