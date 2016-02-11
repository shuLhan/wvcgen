[![GoDoc](https://godoc.org/github.com/shuLhan/wvcgen?status.svg)](https://godoc.org/github.com/shuLhan/wvcgen)

# wvcgen

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

#### Creating Unified Dataset

* Go to directory `merge/wvc2010`
* Run `main.go` script to merge and create new dataset
  ```
  $ go run main.go
  ```

  which will create file `merge_edits_golds.dat` that combine file
  `pan-wvc-2010/edits.csv` with `pan-wvc-2010/gold-annotations.csv` and add
  two new fields. List of files in unified dataset are,

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

One can customize the output of dataset by editing the `merge_edits_gold.dsv`
configuration and run the merge script again.

## Generating Features

After one of PAN WVC dataset has been merged, one can compute the vandalism
features by runnning `main.go` script in root of repository.

    $ go run main.go

Feature values will be written to file `features.dat`.

One can customize the input and which features should be computed by editing
file `features.dsv`,
* `Input` key point to the input file,
* `InputMetadata` contains fields in input file,
* `Output` key point the file where result of features computation will be
  written,
* `OutputMetadata` contain list of features that will computed.
