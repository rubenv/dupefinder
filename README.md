# dupefinder

[![Build Status](https://travis-ci.org/rubenv/dupefinder.svg?branch=master)](https://travis-ci.org/rubenv/dupefinder) [![GoDoc](https://godoc.org/github.com/rubenv/dupefinder?status.png)](https://godoc.org/github.com/rubenv/dupefinder)

Detect duplicate files across different machines, using SHA256

# Installation
```
go get github.com/rubenv/dupefinder/...
```

# Usage

```
Usage: dupefinder -generate filename folder...
    Generates a catalog file at filename based on one or more folders

Usage: dupefinder -detect [-dryrun / -rm] filename folder...
    Detects duplicates using a catalog file in on one or more folders

  -detect=false: Detect duplicate files using a catalog
  -dryrun=false: Print what would be deleted
  -generate=false: Generate a catalog file
  -rm=false: Delete detected duplicates (at your own risk!)
```

# As a library

Import into your application with:

```go
import "github.com/rubenv/dupefinder"
```

## Usage

#### func  Detect

```go
func Detect(catalog string, echo, rm bool, folders ...string) error
```
Detect duplicates. Set echo to true to print duplicates, rm to delete them.

#### func  Generate

```go
func Generate(catalog string, folders ...string) error
```
Generate a catalog file based on a set of folders

#### type DupeCatalog

```go
type DupeCatalog map[string]string
```

Catalog of hash to filename mappings

#### func  ParseCatalog

```go
func ParseCatalog(filename string) (DupeCatalog, error)
```
Parse the catalog file at filename

#### func  ParseCatalogReader

```go
func ParseCatalogReader(reader io.Reader) (DupeCatalog, error)
```
Parse a catalog file using an io.Reader

## License

    (The MIT License)

    Copyright (C) 2015 by Ruben Vermeersch <ruben@rocketeer.be>

    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:

    The above copyright notice and this permission notice shall be included in
    all copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
    THE SOFTWARE.
