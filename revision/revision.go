// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package revision provide the module for working with dump files of Wikipedia
revision.
*/
package revision

import (
	"errors"
	"github.com/shuLhan/tabula"
	"github.com/shuLhan/tekstus/diff"
	"io/ioutil"
	"os"
)

var (
	// Dir define where the revision directory located.
	Dir = ""
	// CleanDir define directory where revision that has been cleaned
	// up located.
	CleanDir = ""
)

/*
SetDir will set revision directory to `path`.
*/
func SetDir(path string) {
	Dir = path
}

/*
SetCleanDir set directory where revision that has been cleaned up located.
*/
func SetCleanDir(path string) {
	CleanDir = path
}

/*
GetContent will return content of revision based on specific `id`.
*/
func GetContent(id string) (string, error) {
	if Dir == "" {
		return "", errors.New("Revision directory is not set!")
	}

	path := Dir + "/" + id + ".txt"

	b, e := ioutil.ReadFile(path)

	return string(b), e
}

/*
GetContentClean return content of revision that has been cleaning up.
*/
func GetContentClean(id string) (string, error) {
	if CleanDir == "" {
		return "", errors.New("Clean revision directory is not set!")
	}

	path := CleanDir + "/" + id + ".txt"

	b, e := ioutil.ReadFile(path)

	return string(b), e
}

/*
GetSize return the file size of revision file.
*/
func GetSize(id string) int64 {
	if Dir == "" {
		return 0
	}

	path := Dir + "/" + id + ".txt"

	finfo, e := os.Stat(path)
	if e != nil {
		return 0
	}

	return finfo.Size()
}

/*
Diff given two list of old and new revision id, compare each of them and save
their diff (deletions and additions as single string joined by " ") on dataset.

The revision text will be looked up in directory `Dir` and with extension
".txt".
*/
func Diff(oldids, newids []string, ext string) (
	tabula.DatasetInterface,
	error,
) {
	diffset := &tabula.Dataset{
		Mode: tabula.DatasetModeColumns,
	}

	colAdds := tabula.NewColumn(tabula.TString, "additions")
	colDels := tabula.NewColumn(tabula.TString, "deletions")

	// Get minimum length
	minlen := len(oldids)
	newlen := len(newids)
	if newlen < minlen {
		minlen = newlen
	}

	for x := 0; x < minlen; x++ {
		oldrevid := Dir + "/" + oldids[x] + ext
		newrevid := Dir + "/" + newids[x] + ext

		diffs, e := diff.Files(oldrevid, newrevid, diff.LevelWords)
		if e != nil {
			return nil, e
		}

		dels := diffs.GetAllDels()
		delstr := dels.Join(" ")
		delrec, e := tabula.NewRecordBy(delstr, tabula.TString)

		if e != nil {
			return nil, e
		}

		adds := diffs.GetAllAdds()
		addstr := adds.Join(" ")
		addrec, e := tabula.NewRecordBy(addstr, tabula.TString)

		if e != nil {
			return nil, e
		}

		colDels.PushBack(delrec)
		colAdds.PushBack(addrec)
	}

	diffset.PushColumn(*colDels)
	diffset.PushColumn(*colAdds)

	return diffset, nil
}
