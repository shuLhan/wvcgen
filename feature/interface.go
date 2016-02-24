// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/tabula"
)

/*
Interface define the methods that must be implemented by feature.
*/
type Interface interface {
	tabula.ColumnInterface
	Compute(dataset tabula.Dataset)
}

/*
ListFeature is a global variables which contain all implemented features.
*/
var ListFeature []Interface

/*
GetByName return feature object by their name.
*/
func GetByName(name string) Interface {
	for _, ftr := range ListFeature {
		if name == ftr.GetName() {
			return ftr
		}
	}
	return nil
}

/*
Register a feature to the list of global features.
*/
func Register(ftr Interface, tipe int, name string) {
	ftr.SetType(tipe)
	ftr.SetName(name)
	ListFeature = append(ListFeature, ftr)
}
