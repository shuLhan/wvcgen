# wvc-generator

This is Wikipedia vandalism dataset generator.
This repository does not provide the full Wikipedia vandalism dataset provided
by uni-weimar.de but provide the script to work with dataset, for example
diff-ing revision, creating new dataset from it, and computing the features.

The generator is written using [Go lang](https://golang.org).

## How To Use

### PAN-WVC-2010

* Download the full dataset from [uni-weimar.de site](http://www.uni-weimar.de/medien/webis/corpora/corpus-pan-wvc-10/pan-wikipedia-vandalism-corpus-2010.zip)
* Extract the zip file
* Rename the extracted directory from `pan-wikipedia-vandalism-corpus-2010` to
  `pan-wvc-2010`
* Move all files in `pan-wvc-2010/article-revisions/partXX/` to
  `pan-wvc-2010/revisions`
* Run `pan-wvc-2010_merge_edits_gold.go` script to create diff of dataset
  ```
  $ go run pan-wvc-2010_merge_edits_gold.go
  ```

  which will create file `pan-wvc-2010_merge_edits_golds.data` that combine
  file `pan-wvc-2010/edits.csv` with `pan-wvc-2010/gold-annotations.csv` and
  add words diff in it,

  * editid
  * class
  * oldrevisionid
  * newrevisionid
  * edittime
  * editor
  * articletitle
  * editcomment
  * deletions
  * additions

The new fields are `deletions` and `additions` which contain diff of old
revision with new revision at words level.

One can customize the output of dataset by editing the
`pan-wvc-2010_merge_edits_gold.dsv` configuration and run the merge script
again.
