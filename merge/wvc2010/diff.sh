#!/bin/sh

DIR_REVISIONS=../../pan-wvc-2010/revisions

diff -u ${DIR_REVISIONS}/${1}.txt ${DIR_REVISIONS}/${2}.txt
