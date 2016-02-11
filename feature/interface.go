// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/dsv"
)

/*
Interface define methods for feature.
*/
type Interface interface {
	GetType() int
	GetName() string
	GetValues() dsv.Column
	Compute(dsv.Dataset)
}

/*
ListFeature is a global variables which contain all implemented features.
*/
var ListFeature []Interface

/*
ListFeatureAdd will add new feature to the list.
*/
func ListFeatureAdd(feature Interface) {
	ListFeature = append(ListFeature, feature)
}

/*
ListFeatureGetByName return feature object by their name.
*/
func ListFeatureGetByName(name string) Interface {
	for _, ftr := range ListFeature {
		if name == ftr.GetName() {
			return ftr
		}
	}
	return nil
}
