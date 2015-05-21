# dupefinder

[![Build Status](https://travis-ci.org/rubenv/dupefinder.svg?branch=master)](https://travis-ci.org/rubenv/dupefinder) [![GoDoc](https://godoc.org/github.com/rubenv/dupefinder?status.png)](https://godoc.org/github.com/rubenv/dupefinder)

Detect duplicate files across different machines

## Installation
```
go get github.com/rubenv/dupefinder/...
```

## Usage as a library

Import into your application with:

```go
import "github.com/rubenv/dupefinder"
```

## Usage

#### func  Detect

```go
func Detect(catalog string, echo, rm bool, folders ...string) error
```

#### func  Generate

```go
func Generate(catalog string, folders ...string) error
```

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
