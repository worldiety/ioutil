# ioutil [![Travis-CI](https://travis-ci.com/worldiety/ioutil.svg?branch=master)](https://travis-ci.com/worldiety/ioutil) [![Go Report Card](https://goreportcard.com/badge/github.com/worldiety/ioutil)](https://goreportcard.com/report/github.com/worldiety/ioutil) [![GoDoc](https://godoc.org/github.com/worldiety/ioutil?status.svg)](http://godoc.org/github.com/worldiety/ioutil)

The package `github.com/worldiety/ioutil` provides a bunch of io helpers, which shifts the "clearness" of some 
standard library design decisions towards a more opinionated, shorter and more comfortable code.

Summary:
* Contains DataOutput and DataInput interfaces and implementations to conveniently work with byte order specific
serialization of numbers and byte slices.
* Provides support for reading and writing 24-, 40-, 48- and 56-bit uint and int support. 
* a more fluent file API  

